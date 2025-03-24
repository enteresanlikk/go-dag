package nodesSlack

import (
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

// type
type SlackNode struct {
	nodesCommon.BaseNode

	Settings map[string]interface{}
}

// base node settings
var baseNode = nodesCommon.BaseNode{
	ID:   "slack",
	Name: "Slack",
}

// create
func NewSlackNode(settings map[string]interface{}) *SlackNode {
	return &SlackNode{
		BaseNode: baseNode,
		Settings: settings,
	}
}

// execute
func (n *SlackNode) Execute(inputs []interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) []interface{} {
		time.Sleep(3 * time.Second)

		image := inputs[0].(string)

		//business logic
		message := "New AI-generated image saved at: " + image

		return []interface{}{message}
	})
}

// factory
func CreateSlackNode(settings map[string]interface{}) (nodesCommon.Node, error) {
	return NewSlackNode(settings), nil
}

// init
func init() {
	factory := nodesCommon.GetFactory()
	factory.Register(baseNode.ID, CreateSlackNode)
}
