package node

import (
	"sync"
)

type Edge struct {
	TargetNode  *Node
	OutputIndex int
}

type Node struct {
	ID       string
	Name     string
	Parents  []*Node
	Children []Edge
	Process  func(inputs []interface{}) []interface{}
	Output   []interface{}
	Mutex    sync.Mutex
	Done     bool
	Settings map[string]interface{}
}

type NodeManager struct {
	nodes map[string]*Node
	mutex sync.RWMutex
}

func GetInstance() *NodeManager {
	return &NodeManager{
		nodes: make(map[string]*Node),
	}
}

func (nm *NodeManager) CreateNode(id, name string, process func(inputs []interface{}) []interface{}, opts ...Option) *Node {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	if node, exists := nm.nodes[id]; exists {
		return node
	}

	node := &Node{
		ID:       id,
		Name:     name,
		Process:  process,
		Settings: make(map[string]interface{}),
		Children: make([]Edge, 0),
	}

	baseNode := NewBaseNode(id, name, opts...)
	for k, v := range baseNode.Settings {
		node.Settings[k] = v
	}

	nm.nodes[id] = node
	return node
}

func (nm *NodeManager) GetNode(id string) (*Node, bool) {
	nm.mutex.RLock()
	defer nm.mutex.RUnlock()

	node, exists := nm.nodes[id]
	return node, exists
}

func (nm *NodeManager) AddEdge(parentID, childID string, outputIndex int) bool {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	parentNode, existsParent := nm.nodes[parentID]
	childNode, existsChild := nm.nodes[childID]

	if !existsParent || !existsChild {
		return false
	}

	edge := Edge{
		TargetNode:  childNode,
		OutputIndex: outputIndex,
	}

	parentNode.Children = append(parentNode.Children, edge)
	childNode.Parents = append(childNode.Parents, parentNode)
	return true
}

func (n *Node) GetSetting(key string, defaultValue interface{}) interface{} {
	if value, exists := n.Settings[key]; exists {
		return value
	}
	return defaultValue
}

func (nm *NodeManager) AddNode(node *Node) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	if _, exists := nm.nodes[node.ID]; !exists {
		nm.nodes[node.ID] = node
	}
}
