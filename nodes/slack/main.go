package nodesSlack

import (
	"fmt"
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

type SlackNode struct {
	nodesCommon.BaseNode

	Webhook string
}

func NewSlackNode(webhook string) *SlackNode {
	return &SlackNode{
		BaseNode: nodesCommon.BaseNode{
			Enabled: true,
			Name:    "Slack",
		},
		Webhook: webhook,
	}
}

func (n *SlackNode) Execute(input interface{}) interface{} {
	return n.ExecuteWithCheck(input, func(input interface{}) interface{} {
		time.Sleep(3 * time.Second)

		message := fmt.Sprintf("New AI-generated image saved at: %s", input.(string))
		fmt.Printf("Sending Slack Notification to %s: %s\n", n.Webhook, message)

		if n.Next != nil {
			return n.Next.Execute(input)
		}
		return input
	})
}
