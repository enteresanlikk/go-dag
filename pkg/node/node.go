package node

import (
	"sync"
)

type Edge struct {
	TargetNode *Node
	OutputKey  string
}

type Node struct {
	ID       string
	Name     string
	Parents  []*Node
	Children []Edge
	Process  func(inputs map[string]interface{}) map[string]interface{}
	Output   map[string]interface{}
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

func (nm *NodeManager) CreateNode(id, name string, process func(inputs map[string]interface{}) map[string]interface{}) *Node {
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

	baseNode := NewBaseNode(id, name)
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

func (nm *NodeManager) AddEdge(parentID, childID, outputKey string) bool {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	parentNode, existsParent := nm.nodes[parentID]
	childNode, existsChild := nm.nodes[childID]

	if !existsParent || !existsChild {
		return false
	}

	edge := Edge{
		TargetNode: childNode,
		OutputKey:  "",
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
