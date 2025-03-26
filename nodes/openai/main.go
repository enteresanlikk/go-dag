package nodesOpenAI

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type OpenAINode struct {
	node.BaseNode
}

func NewOpenAINode() *OpenAINode {
	return &OpenAINode{
		BaseNode: node.NewBaseNode("openai", "OpenAI"),
	}
}

func (n *OpenAINode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("OpenAINode Process", inputs)

	prompt := inputs["prompt"].(string)

	response := "OpenAI Response for: " + prompt

	return map[string]interface{}{
		"response": response,
	}
}

func init() {
	node.RegisterProcessor(NewOpenAINode())
}
