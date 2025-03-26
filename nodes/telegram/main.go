package nodesTelegram

import (
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

// type
type TelegramNode struct {
	nodesCommon.BaseNode

	Settings map[string]interface{}
}

// base node settings
var baseNode = nodesCommon.BaseNode{
	ID:   "telegram",
	Name: "Telegram",
}

// create
func NewTelegramNode(settings map[string]interface{}) *TelegramNode {
	return &TelegramNode{
		BaseNode: baseNode,
		Settings: settings,
	}
}

// execute
func (n *TelegramNode) Execute(inputs []interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) []interface{} {
		// time.Sleep(1 * time.Second)

		image := inputs[0].(string)

		//business logic
		message := "New AI-generated image saved at: " + image

		return []interface{}{message}
	})
}

// factory
func CreateTelegramNode(settings map[string]interface{}) (nodesCommon.Node, error) {
	return NewTelegramNode(settings), nil
}

// init
func init() {
	factory := nodesCommon.GetFactory()
	factory.Register(baseNode.ID, CreateTelegramNode)
}
