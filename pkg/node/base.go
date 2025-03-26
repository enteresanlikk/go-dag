package node

type BaseNode struct {
	id       string
	name     string
	Settings map[string]interface{}
}

func NewBaseNode(id, name string, opts ...Option) BaseNode {
	node := BaseNode{
		id:       id,
		name:     name,
		Settings: make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(&node)
	}

	return node
}

func (n *BaseNode) ID() string {
	return n.id
}

func (n *BaseNode) Name() string {
	return n.name
}

func (n *BaseNode) SetSetting(key string, value interface{}) {
	n.Settings[key] = value
}

func (n *BaseNode) GetSetting(key string, defaultValue interface{}) interface{} {
	if value, exists := n.Settings[key]; exists {
		return value
	}
	return defaultValue
}

type Option func(node *BaseNode)

func WithSetting(key string, value interface{}) Option {
	return func(node *BaseNode) {
		node.SetSetting(key, value)
	}
}
