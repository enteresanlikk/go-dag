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

func (n *SlackNode) Execute(inputs ...interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) interface{} {
		time.Sleep(3 * time.Second)

		message := fmt.Sprintf("New AI-generated image saved at: %s", inputs[0].(string))
		fmt.Printf("Sending Slack Notification to %s: %s\n", n.Webhook, message)

		return message
	})
}
