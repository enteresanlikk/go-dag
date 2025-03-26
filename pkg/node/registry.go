package node

type NodeProcessor interface {
	GetID() string
	GetName() string
	Process(inputs map[string]interface{}) map[string]interface{}
	GetSetting(key string, defaultValue interface{}) interface{}
	SetSetting(key string, value interface{})
	GetSettings() map[string]interface{}
	SetSettings(settings map[string]interface{})
}

var (
	nodeRegistry = make(map[string]NodeProcessor)
)

func RegisterProcessor(processor NodeProcessor) {
	nodeRegistry[processor.GetID()] = processor
}

func GetProcessor(id string) (NodeProcessor, bool) {
	processor, exists := nodeRegistry[id]
	return processor, exists
}

func InitializeNodes(manager *NodeManager) {
	for _, processor := range nodeRegistry {
		manager.CreateNode(processor.GetID(), processor.GetName(), processor.Process)
	}
}
