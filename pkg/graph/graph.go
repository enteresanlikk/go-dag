package graph

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// Graph represents a DAG
type Graph struct {
	nodeManager *node.NodeManager
	wg          sync.WaitGroup
}

// NewGraph creates a new DAG
func NewGraph() *Graph {
	return &Graph{
		nodeManager: node.GetInstance(),
	}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(id, name string, process func(inputs []interface{}) []interface{}) {
	g.nodeManager.CreateNode(id, name, process)
}

// AddEdge creates a connection between nodes in the graph
func (g *Graph) AddEdge(parentID, childID string) {
	g.nodeManager.AddEdge(parentID, childID)
}

// Execute runs the DAG starting from the specified node
func (g *Graph) Execute(startNodeID string, inputs []interface{}) {
	startNode, exists := g.nodeManager.GetNode(startNodeID)
	if !exists {
		return
	}

	g.wg.Add(1)
	go g.executeNode(startNode, inputs)
	g.wg.Wait()
}

// executeNode runs a specific node and propagates data to connected nodes
func (g *Graph) executeNode(node *node.Node, inputs []interface{}) {
	defer g.wg.Done()

	// Execute node's process function
	node.Mutex.Lock()
	if !node.Done {
		node.Output = node.Process(inputs)
		node.Done = true
	}
	node.Mutex.Unlock()

	// Execute child nodes
	for _, child := range node.Children {
		allParentsReady := true
		for _, parent := range child.Parents {
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
			go g.executeNode(child, g.collectInputs(child))
		}
	}
}

// collectInputs gathers all parent nodes' outputs
func (g *Graph) collectInputs(node *node.Node) []interface{} {
	inputs := []interface{}{}
	for _, parent := range node.Parents {
		inputs = append(inputs, parent.Output...)
	}
	return inputs
}

// GetNodeManager returns the node manager instance
func (g *Graph) GetNodeManager() *node.NodeManager {
	return g.nodeManager
}

// LoadFromJSON loads graph configuration from JSON
func (g *Graph) LoadFromJSON(jsonData *GraphConfig) error {
	// Initialize nodes with settings
	for _, nodeConfig := range jsonData.Nodes {
		processor, exists := node.GetProcessor(nodeConfig.ID)
		if !exists {
			return fmt.Errorf("node processor not found for ID: %s", nodeConfig.ID)
		}

		// Convert and apply settings
		for key, value := range nodeConfig.Settings {
			if num, ok := value.(json.Number); ok {
				if n, err := num.Int64(); err == nil {
					processor.SetSetting(key, int(n))
					continue
				}
			}
			processor.SetSetting(key, value)
		}

		// Create node with settings
		node := &node.Node{
			ID:       processor.ID(),
			Name:     processor.Name(),
			Process:  processor.Process,
			Settings: make(map[string]interface{}),
		}

		// Copy settings from processor to node
		for key, value := range nodeConfig.Settings {
			if num, ok := value.(json.Number); ok {
				if n, err := num.Int64(); err == nil {
					node.Settings[key] = int(n)
					continue
				}
			}
			node.Settings[key] = value
		}

		g.nodeManager.AddNode(node)

		// Process inputs if provided
		if len(nodeConfig.Inputs) > 0 {
			inputs := make([]interface{}, 0, len(nodeConfig.Inputs))
			inputs = append(inputs, nodeConfig.Inputs...)
			if len(inputs) > 0 {
				g.wg.Add(1)
				go g.executeNode(node, inputs)
			}
		}
	}

	// Add edges
	for _, edge := range jsonData.Edges {
		g.AddEdge(edge.Source, edge.Target)
	}

	g.wg.Wait() // Wait for all nodes to complete

	return nil
}
