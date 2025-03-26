package nodesSlack

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type SlackNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newSlackNode() *SlackNode {
	return &SlackNode{
		BaseNode: node.NewBaseNode("slack", "Slack"),
	}
}

func (n *SlackNode) Process(inputs []interface{}) []interface{} {
	image := inputs[0].(string)

	message := "New AI-generated image saved at: " + image

	return []interface{}{message}
}

func init() {
	node.RegisterProcessor(newSlackNode())
}
