package graph

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type Graph struct {
	nodeManager *node.NodeManager
	wg          sync.WaitGroup
	nodeOutputs map[string]map[string]interface{}
}

func NewGraph() *Graph {
	return &Graph{
		nodeManager: node.GetInstance(),
		nodeOutputs: make(map[string]map[string]interface{}),
	}
}

func (g *Graph) AddNode(id, name string, process func(map[string]interface{}) map[string]interface{}) {
	g.nodeManager.CreateNode(id, name, process)
}

func (g *Graph) AddEdge(parentID, childID string, outputKey int) {
	g.nodeManager.AddEdge(parentID, childID, outputKey)
}

func (g *Graph) Execute(startNodeID string, inputs map[string]interface{}) {
	startNode, exists := g.nodeManager.GetNode(startNodeID)
	if !exists {
		return
	}

	g.wg.Add(1)
	go g.executeNode(startNode, inputs)
	g.wg.Wait()
}

func (g *Graph) executeNode(node *node.Node, inputs map[string]interface{}) {
	defer g.wg.Done()

	node.Mutex.Lock()
	if !node.Done {
		node.Output = node.Process(inputs)
		node.Done = true
		g.nodeOutputs[node.ID] = node.Output
	}
	outputs := node.Output
	node.Mutex.Unlock()

	processedEdges := make(map[string]bool)

	for _, edge := range node.Children {
		if processedEdges[edge.OutputKey] || outputs[edge.OutputKey] == nil {
			continue
		}

		processedEdges[edge.OutputKey] = true

		allParentsReady := true
		for _, parent := range edge.TargetNode.Parents {
			parent.Mutex.Lock()
			if !parent.Done {
				allParentsReady = false
			}
			parent.Mutex.Unlock()

			if !allParentsReady {
				break
			}
		}

		if allParentsReady {
			g.wg.Add(1)
			targetInputs := make(map[string]interface{})
			targetInputs[edge.OutputKey] = outputs[edge.OutputKey]
			go g.executeNode(edge.TargetNode, targetInputs)
		}
	}
}

func (g *Graph) GetNodeManager() *node.NodeManager {
	return g.nodeManager
}

func (g *Graph) processInputValue(value interface{}, nodeID string) (interface{}, error) {
	if strValue, ok := value.(string); ok {
		// Check if the value contains a dynamic reference
		re := regexp.MustCompile(`\$id\[([^\]]+)\]\.outputs\[([^\]]+)\]`)
		matches := re.FindStringSubmatch(strValue)

		if len(matches) == 3 {
			// Extract referenced node ID and output key
			referencedNodeID := matches[1]
			outputKey := matches[2]

			// Get the referenced node's outputs
			if outputs, exists := g.nodeOutputs[referencedNodeID]; exists {
				if output, exists := outputs[outputKey]; exists {
					// If the entire string is just the reference, return the output value directly
					if matches[0] == strValue {
						return output, nil
					}
					// Otherwise, replace the reference in the string with the output value
					return strings.Replace(strValue, matches[0], fmt.Sprintf("%v", output), -1), nil
				}
			}
			return nil, fmt.Errorf("referenced output not found: node=%s, key=%s", referencedNodeID, outputKey)
		}
	}
	return value, nil
}

type dependencyGraph struct {
	nodes map[string]NodeConfig
	edges map[string][]string // node -> dependencies
}

func (g *Graph) buildDependencyGraph(jsonData *GraphConfig) *dependencyGraph {
	graph := &dependencyGraph{
		nodes: make(map[string]NodeConfig),
		edges: make(map[string][]string),
	}

	// Add all nodes to the graph
	for _, node := range jsonData.Nodes {
		graph.nodes[node.ID] = node
		graph.edges[node.ID] = make([]string, 0)
	}

	// Add edges from the edges configuration
	for _, edge := range jsonData.Edges {
		graph.edges[edge.Source] = append(graph.edges[edge.Source], edge.Target)
	}

	return graph
}

func (g *Graph) topologicalSort(jsonData *GraphConfig) ([]NodeConfig, error) {
	graph := g.buildDependencyGraph(jsonData)
	visited := make(map[string]bool)
	temp := make(map[string]bool)
	order := make([]NodeConfig, 0)

	var visit func(string) error
	visit = func(nodeID string) error {
		if temp[nodeID] {
			return fmt.Errorf("cycle detected in workflow")
		}
		if visited[nodeID] {
			return nil
		}
		temp[nodeID] = true

		// Visit all dependencies
		for _, dep := range graph.edges[nodeID] {
			if err := visit(dep); err != nil {
				return err
			}
		}

		temp[nodeID] = false
		visited[nodeID] = true
		order = append([]NodeConfig{graph.nodes[nodeID]}, order...)
		return nil
	}

	// Visit all nodes
	for nodeID := range graph.nodes {
		if !visited[nodeID] {
			if err := visit(nodeID); err != nil {
				return nil, err
			}
		}
	}

	return order, nil
}

func (g *Graph) LoadFromJSON(jsonData *GraphConfig) error {
	// First pass: Create all nodes
	for _, nodeConfig := range jsonData.Nodes {
		processor, exists := node.GetProcessor(nodeConfig.ID)
		if !exists {
			return fmt.Errorf("node processor not found for ID: %s", nodeConfig.ID)
		}

		processor.SetSettings(nodeConfig.Settings)

		node := &node.Node{
			ID:       processor.GetID(),
			Name:     processor.GetName(),
			Process:  processor.Process,
			Settings: make(map[string]interface{}),
			Children: make([]node.Edge, 0),
		}

		for key, value := range nodeConfig.Settings {
			node.Settings[key] = value
		}

		g.nodeManager.AddNode(node)
	}

	// Second pass: Create edges
	for _, edge := range jsonData.Edges {
		g.AddEdge(edge.Source, edge.Target, edge.OutputKey)
	}

	// Third pass: Sort nodes topologically and execute them in order
	sortedNodes, err := g.topologicalSort(jsonData)
	if err != nil {
		return err
	}

	// Execute nodes in topological order
	for _, nodeConfig := range sortedNodes {
		node, _ := g.nodeManager.GetNode(nodeConfig.ID)
		inputs := make(map[string]interface{})

		if len(nodeConfig.Inputs) > 0 {
			// Process each input
			for key, value := range nodeConfig.Inputs {
				processedValue, err := g.processInputValue(value, nodeConfig.ID)
				if err != nil {
					return err
				}
				inputs[key] = processedValue
			}
		}

		// Execute node even if it has no inputs
		g.wg.Add(1)
		go g.executeNode(node, inputs)
		g.wg.Wait() // Wait for this node to complete before processing next node
	}

	return nil
}
