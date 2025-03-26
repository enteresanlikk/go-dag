package nodesSlack

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type SlackNode struct {
	node.BaseNode
}

func NewSlackNode() *SlackNode {
	return &SlackNode{
		BaseNode: node.NewBaseNode("slack", "Slack"),
	}
}

func (n *SlackNode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("SlackNode Process", inputs)

	message := inputs["message"].(string)

	return map[string]interface{}{
		"message": message,
	}
}

func init() {
	node.RegisterProcessor(NewSlackNode())
}
