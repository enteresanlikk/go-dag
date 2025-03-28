package node

import (
	"sync"
)

type NodeProcessor interface {
	GetID() string
	GetName() string
	Process(inputs map[string]interface{}) map[string]interface{}
	GetSetting(key string, defaultValue interface{}) interface{}
	SetSetting(key string, value interface{})
	GetSettings() map[string]interface{}
	SetSettings(settings map[string]interface{})
}

type NodeRegistry struct {
	processors sync.Map
	cache      sync.Map
}

var (
	instance *NodeRegistry
	once     sync.Once
)

func GetRegistry() *NodeRegistry {
	once.Do(func() {
		instance = &NodeRegistry{}
	})
	return instance
}

func (r *NodeRegistry) RegisterProcessor(processor NodeProcessor) {
	r.processors.Store(processor.GetID(), processor)
}

func (r *NodeRegistry) GetProcessor(id string) (NodeProcessor, bool) {
	if cached, ok := r.cache.Load(id); ok {
		return cached.(NodeProcessor), true
	}

	if value, exists := r.processors.Load(id); exists {
		processor := value.(NodeProcessor)
		r.cache.Store(id, processor)
		return processor, true
	}
	return nil, false
}

func RegisterProcessor(processor NodeProcessor) {
	GetRegistry().RegisterProcessor(processor)
}

func GetProcessor(id string) (NodeProcessor, bool) {
	return GetRegistry().GetProcessor(id)
}
