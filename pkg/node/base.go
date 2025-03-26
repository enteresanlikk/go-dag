package node

// BaseNode provides common functionality for all nodes
type BaseNode struct {
	id       string
	name     string
	Settings map[string]interface{}
}

// NewBaseNode creates a new base node
func NewBaseNode(id, name string, opts ...Option) BaseNode {
	node := BaseNode{
		id:       id,
		name:     name,
		Settings: make(map[string]interface{}),
	}

	// Apply options
	for _, opt := range opts {
		opt(&node)
	}

	return node
}

// ID returns the node's ID
func (n *BaseNode) ID() string {
	return n.id
}

// Name returns the node's name
func (n *BaseNode) Name() string {
	return n.name
}

// SetSetting sets a setting value
func (n *BaseNode) SetSetting(key string, value interface{}) {
	n.Settings[key] = value
}

// GetSetting gets a setting value with a default value if not found
func (n *BaseNode) GetSetting(key string, defaultValue interface{}) interface{} {
	if value, exists := n.Settings[key]; exists {
		return value
	}
	return defaultValue
}

// Option represents a node option
type Option func(node *BaseNode)

// WithSetting creates an option to set a setting
func WithSetting(key string, value interface{}) Option {
	return func(node *BaseNode) {
		node.SetSetting(key, value)
	}
}
