package nodesGemini

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type GeminiNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newGeminiNode() *GeminiNode {
	return &GeminiNode{
		BaseNode: node.NewBaseNode("gemini", "Gemini"),
	}
}

func (n *GeminiNode) Process(inputs []interface{}) []interface{} {
	prompt := inputs[0].(string)

	return []interface{}{prompt}
}

func init() {
	node.RegisterProcessor(newGeminiNode())
}
