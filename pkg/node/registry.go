package node

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

func RegisterProcessor(processor NodeProcessor) {
	nodeRegistry[processor.ID()] = processor
}

func GetProcessor(id string) (NodeProcessor, bool) {
	processor, exists := nodeRegistry[id]
	return processor, exists
}

func InitializeNodes(manager *NodeManager) {
	for _, processor := range nodeRegistry {
		manager.CreateNode(processor.ID(), processor.Name(), processor.Process)
	}
}
