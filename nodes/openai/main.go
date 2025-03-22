package nodesOpenAI

import (
	"fmt"
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

type OpenAINode struct {
	nodesCommon.BaseNode

	ApiKey string
}

func NewOpenAINode(apiKey string) *OpenAINode {
	return &OpenAINode{
		BaseNode: nodesCommon.BaseNode{
			Enabled: true,
			Name:    "OpenAI",
		},
		ApiKey: apiKey,
	}
}

func (n *OpenAINode) Execute(inputs ...interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) interface{} {
		time.Sleep(5 * time.Second)

		prompt := inputs[0].(string)
		response := fmt.Sprintf("OpenAI Response for: %s", prompt)
		fmt.Println("OpenAI Output:", response)

		return response
	})
}
