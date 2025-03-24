package nodesCommon

import "errors"

type NodeCreator func(settings map[string]interface{}) (Node, error)

var defaultFactory *NodeFactory

type NodeFactory struct {
	creators map[string]NodeCreator
}

func init() {
	defaultFactory = &NodeFactory{
		creators: make(map[string]NodeCreator),
	}
}

func GetFactory() *NodeFactory {
	return defaultFactory
}

func (f *NodeFactory) Register(nodeType string, creator NodeCreator) {
	f.creators[nodeType] = creator
}

func (f *NodeFactory) Create(nodeType string, settings map[string]interface{}) (Node, error) {
	creator, exists := f.creators[nodeType]
	if !exists {
		return nil, errors.New("unknown node type: " + nodeType)
	}
	return creator(settings)
}

func (f *NodeFactory) GetAllNodes() map[string]NodeCreator {
	return f.creators
}

func (f *NodeFactory) GetNode(nodeType string) NodeCreator {
	return f.creators[nodeType]
}
