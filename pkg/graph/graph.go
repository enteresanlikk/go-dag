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
	nodeOutputs sync.Map
}

func NewGraph() *Graph {
	return &Graph{
		nodeManager: node.GetInstance(),
		wg:          sync.WaitGroup{},
		nodeOutputs: sync.Map{},
	}
}

func (g *Graph) AddNode(id, name string, process func(map[string]interface{}) map[string]interface{}) {
	g.nodeManager.CreateNode(id, name, process)
}

func (g *Graph) AddEdge(parentID, childID, outputKey string) {
	g.nodeManager.AddEdge(parentID, childID, outputKey)
}

func (g *Graph) Execute(startNodeID string, inputs map[string]interface{}) {
	startNode, exists := g.nodeManager.GetNode(startNodeID)
	if !exists {
		return
	}

	g.wg.Add(1)
	g.ExecuteNode(startNode, inputs)
	g.wg.Wait()
}

func (g *Graph) ExecuteNode(node *node.Node, inputs map[string]interface{}) {
	defer g.wg.Done()

	node.Mutex.Lock()
	if !node.Done {
		output := node.Process(inputs)
		node.Output = output
		node.Done = true
		g.nodeOutputs.Store(node.ID, output)
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
			isDone := parent.Done
			parent.Mutex.Unlock()

			if !isDone {
				allParentsReady = false
				break
			}
		}

		if allParentsReady {
			g.wg.Add(1)
			targetInputs := make(map[string]interface{})
			targetInputs[edge.OutputKey] = outputs[edge.OutputKey]
			g.ExecuteNode(edge.TargetNode, targetInputs)
		}
	}
}

func (g *Graph) GetNodeManager() *node.NodeManager {
	return g.nodeManager
}

var reCache = regexp.MustCompile(`\$([^.]+)\.([^.]+)`)

func (g *Graph) ProcessInputValue(value interface{}) (interface{}, error) {
	if strValue, ok := value.(string); ok {
		matches := reCache.FindStringSubmatch(strValue)
		if len(matches) == 3 {
			referencedNodeID := matches[1]
			outputKey := matches[2]

			if outputsRaw, exists := g.nodeOutputs.Load(referencedNodeID); exists {
				outputs := outputsRaw.(map[string]interface{})
				if output, exists := outputs[outputKey]; exists {
					if matches[0] == strValue {
						return output, nil
					}
					return strings.Replace(strValue, matches[0], fmt.Sprintf("%v", output), -1), nil
				}
			}
			return nil, fmt.Errorf("referenced output not found: node=%s, key=%s", referencedNodeID, outputKey)
		}
	}
	return value, nil
}

type DependencyGraph struct {
	nodes map[string]NodeConfig
	edges map[string][]string
}

func (g *Graph) BuildDependencyGraph(jsonData *GraphConfig) *DependencyGraph {
	graph := &DependencyGraph{
		nodes: make(map[string]NodeConfig),
		edges: make(map[string][]string),
	}

	for _, node := range jsonData.Nodes {
		graph.nodes[node.ID] = node
		graph.edges[node.ID] = make([]string, 0)
	}

	for _, edge := range jsonData.Edges {
		graph.edges[edge.Source] = append(graph.edges[edge.Source], edge.Target)
	}

	return graph
}

func (g *Graph) TopologicalSort(jsonData *GraphConfig) ([]NodeConfig, error) {
	graph := g.BuildDependencyGraph(jsonData)
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
	for _, nodeConfig := range jsonData.Nodes {
		processor, exists := node.GetProcessor(nodeConfig.ID)
		if !exists {
			return fmt.Errorf("node processor not found for ID: %s", nodeConfig.ID)
		}

		node := &node.Node{
			ID:       processor.GetID(),
			Name:     processor.GetName(),
			Process:  processor.Process,
			Settings: nodeConfig.Settings,
			Children: make([]node.Edge, 0, 4),
		}

		g.nodeManager.AddNode(node)
	}

	for _, edge := range jsonData.Edges {
		g.AddEdge(edge.Source, edge.Target, edge.OutputKey)
	}

	sortedNodes, err := g.TopologicalSort(jsonData)
	if err != nil {
		return err
	}

	for _, nodeConfig := range sortedNodes {
		node, _ := g.nodeManager.GetNode(nodeConfig.ID)
		inputs := make(map[string]interface{})

		if len(nodeConfig.Inputs) > 0 {
			for key, value := range nodeConfig.Inputs {
				processedValue, err := g.ProcessInputValue(value)
				if err != nil {
					return err
				}
				inputs[key] = processedValue
			}
		}

		g.wg.Add(1)
		g.ExecuteNode(node, inputs)
		g.wg.Wait()
	}

	return nil
}
