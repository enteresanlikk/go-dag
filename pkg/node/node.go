package node

import (
	"sync"
)

// Node represents a node in the DAG
type Node struct {
	ID       string
	Name     string
	Parents  []*Node
	Children []*Node
	Process  func(inputs []interface{}) []interface{}
	Output   []interface{}
	Mutex    sync.Mutex
	Done     bool
	Settings map[string]interface{}
}

var (
	nodeInstance *NodeManager
	once         sync.Once
)

// NodeManager manages node operations with singleton pattern
type NodeManager struct {
	nodes map[string]*Node
	mutex sync.RWMutex
}

// GetInstance returns the singleton instance of NodeManager
func GetInstance() *NodeManager {
	once.Do(func() {
		nodeInstance = &NodeManager{
			nodes: make(map[string]*Node),
		}
	})
	return nodeInstance
}

// CreateNode creates a new node and adds it to the manager
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
	}

	// Apply options
	baseNode := NewBaseNode(id, name, opts...)
	for k, v := range baseNode.Settings {
		node.Settings[k] = v
	}

	nm.nodes[id] = node
	return node
}

// GetNode returns a node by its ID
func (nm *NodeManager) GetNode(id string) (*Node, bool) {
	nm.mutex.RLock()
	defer nm.mutex.RUnlock()

	node, exists := nm.nodes[id]
	return node, exists
}

// AddEdge creates a connection between two nodes
func (nm *NodeManager) AddEdge(parentID, childID string) bool {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	parentNode, existsParent := nm.nodes[parentID]
	childNode, existsChild := nm.nodes[childID]

	if !existsParent || !existsChild {
		return false
	}

	parentNode.Children = append(parentNode.Children, childNode)
	childNode.Parents = append(childNode.Parents, parentNode)
	return true
}

// GetSetting gets a setting value with a default value if not found
func (n *Node) GetSetting(key string, defaultValue interface{}) interface{} {
	if value, exists := n.Settings[key]; exists {
		return value
	}
	return defaultValue
}

// AddNode adds a pre-configured node to the manager
func (nm *NodeManager) AddNode(node *Node) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	if _, exists := nm.nodes[node.ID]; !exists {
		nm.nodes[node.ID] = node
	}
}
