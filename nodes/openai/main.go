package nodesOpenAI

import (
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

// type
type OpenAINode struct {
	nodesCommon.BaseNode

	Settings map[string]interface{}
}

// base node settings
var baseNode = nodesCommon.BaseNode{
	ID:   "openai",
	Name: "OpenAI",
}

// create
func NewOpenAINode(settings map[string]interface{}) *OpenAINode {
	return &OpenAINode{
		BaseNode: baseNode,
		Settings: settings,
	}
}

// execute
func (n *OpenAINode) Execute(inputs []interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) []interface{} {
		time.Sleep(5 * time.Second)

		prompt := inputs[0].(string)

		//business logic
		response := "OpenAI Response for: " + prompt

		return []interface{}{response}
	})
}

// factory
func CreateOpenAINode(settings map[string]interface{}) (nodesCommon.Node, error) {
	return NewOpenAINode(settings), nil
}

// init
func init() {
	factory := nodesCommon.GetFactory()
	factory.Register(baseNode.ID, CreateOpenAINode)
}
