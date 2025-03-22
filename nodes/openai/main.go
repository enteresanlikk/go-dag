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

func (n *OpenAINode) Execute(input interface{}) interface{} {
	return n.ExecuteWithCheck(input, func(input interface{}) interface{} {
		time.Sleep(5 * time.Second)

		prompt := input.(string)
		response := fmt.Sprintf("OpenAI Response for: %s", prompt)
		fmt.Println("OpenAI Output:", response)

		if n.Next != nil {
			return n.Next.Execute(response)
		}
		return response
	})
}
