package node

// NodeProcessor represents a node's processing function
type NodeProcessor interface {
	ID() string
	Name() string
	Process(inputs []interface{}) []interface{}
	GetSetting(key string, defaultValue interface{}) interface{}
	SetSetting(key string, value interface{})
}

var (
	nodeRegistry = make(map[string]NodeProcessor)
)

// RegisterProcessor registers a node processor to the registry
func RegisterProcessor(processor NodeProcessor) {
	nodeRegistry[processor.ID()] = processor
}

// GetProcessor returns a node processor by its ID
func GetProcessor(id string) (NodeProcessor, bool) {
	processor, exists := nodeRegistry[id]
	return processor, exists
}

// InitializeNodes initializes all registered nodes in the graph
func InitializeNodes(manager *NodeManager) {
	for _, processor := range nodeRegistry {
		manager.CreateNode(processor.ID(), processor.Name(), processor.Process)
	}
}
