package nodesOpenAI

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type OpenAINode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newOpenAINode() *OpenAINode {
	return &OpenAINode{
		BaseNode: node.NewBaseNode("openai", "OpenAI"),
	}
}

func (n *OpenAINode) Process(inputs []interface{}) []interface{} {
	prompt := inputs[0].(string)

	response := "OpenAI Response for: " + prompt

	return []interface{}{response}
}

func init() {
	node.RegisterProcessor(newOpenAINode())
}
