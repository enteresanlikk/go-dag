package nodesOpenAI

import (
	"time"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type OpenAINode struct {
	node.BaseNode

	Settings map[string]interface{}
}

// create
func newOpenAINode() *OpenAINode {
	return &OpenAINode{
		BaseNode: node.NewBaseNode("openai", "OpenAI"),
	}
}

// execute
func (n *OpenAINode) Process(inputs []interface{}) []interface{} {
	time.Sleep(5 * time.Second)

	prompt := inputs[0].(string)

	//business logic
	response := "OpenAI Response for: " + prompt

	return []interface{}{response}
}

// init
func init() {
	node.RegisterProcessor(newOpenAINode())
}
