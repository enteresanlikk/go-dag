package nodesSlack

import (
	"time"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type SlackNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

// create
func newSlackNode() *SlackNode {
	return &SlackNode{
		BaseNode: node.NewBaseNode("slack", "Slack"),
	}
}

// execute
func (n *SlackNode) Process(inputs []interface{}) []interface{} {
	time.Sleep(3 * time.Second)

	image := inputs[0].(string)

	//business logic
	message := "New AI-generated image saved at: " + image

	return []interface{}{message}
}

// init
func init() {
	node.RegisterProcessor(newSlackNode())
}
