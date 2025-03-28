package node

type BaseNode struct {
	ID       string
	Name     string
	Settings map[string]interface{}
}

func NewBaseNode(id, name string) BaseNode {
	node := BaseNode{
		ID:       id,
		Name:     name,
		Settings: make(map[string]interface{}),
	}

	return node
}

func (n *BaseNode) GetID() string {
	return n.ID
}

func (n *BaseNode) GetName() string {
	return n.Name
}

func (n *BaseNode) GetSettings() map[string]interface{} {
	return n.Settings
}

func (n *BaseNode) SetSettings(settings map[string]interface{}) {
	n.Settings = settings
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
