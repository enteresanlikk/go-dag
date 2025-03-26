package graph

import (
	"fmt"
	"sync"

	"github.com/goccy/go-json"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type Graph struct {
	nodeManager *node.NodeManager
	wg          sync.WaitGroup
}

func NewGraph() *Graph {
	return &Graph{
		nodeManager: node.GetInstance(),
	}
}

func (g *Graph) AddNode(id, name string, process func(inputs []interface{}) []interface{}) {
	g.nodeManager.CreateNode(id, name, process)
}

func (g *Graph) AddEdge(parentID, childID string, outputIndex int) {
	g.nodeManager.AddEdge(parentID, childID, outputIndex)
}

func (g *Graph) Execute(startNodeID string, inputs []interface{}) {
	startNode, exists := g.nodeManager.GetNode(startNodeID)
	if !exists {
		return
	}

	g.wg.Add(1)
	go g.executeNode(startNode, inputs)
	g.wg.Wait()
}

func (g *Graph) executeNode(node *node.Node, inputs []interface{}) {
	defer g.wg.Done()

	node.Mutex.Lock()
	if !node.Done {
		node.Output = node.Process(inputs)
		node.Done = true
	}
	outputs := node.Output
	node.Mutex.Unlock()

	processedEdges := make(map[int]bool)

	for _, edge := range node.Children {
		if processedEdges[edge.OutputIndex] || edge.OutputIndex >= len(outputs) || outputs[edge.OutputIndex] == nil {
			continue
		}

		processedEdges[edge.OutputIndex] = true

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
			go g.executeNode(edge.TargetNode, []interface{}{outputs[edge.OutputIndex]})
		}
	}
}

func (g *Graph) GetNodeManager() *node.NodeManager {
	return g.nodeManager
}

func (g *Graph) LoadFromJSON(jsonData *GraphConfig) error {
	for _, nodeConfig := range jsonData.Nodes {
		processor, exists := node.GetProcessor(nodeConfig.ID)
		if !exists {
			return fmt.Errorf("node processor not found for ID: %s", nodeConfig.ID)
		}

		for key, value := range nodeConfig.Settings {
			if num, ok := value.(json.Number); ok {
				if n, err := num.Int64(); err == nil {
					processor.SetSetting(key, int(n))
					continue
				}
			}
			processor.SetSetting(key, value)
		}

		node := &node.Node{
			ID:       processor.ID(),
			Name:     processor.Name(),
			Process:  processor.Process,
			Settings: make(map[string]interface{}),
			Children: make([]node.Edge, 0),
		}

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

		if len(nodeConfig.Inputs) > 0 {
			inputs := make([]interface{}, 0, len(nodeConfig.Inputs))
			inputs = append(inputs, nodeConfig.Inputs...)
			if len(inputs) > 0 {
				g.wg.Add(1)
				go g.executeNode(node, inputs)
			}
		}
	}

	for _, edge := range jsonData.Edges {
		g.AddEdge(edge.Source, edge.Target, edge.OutputIndex)
	}

	g.wg.Wait()

	return nil
}
