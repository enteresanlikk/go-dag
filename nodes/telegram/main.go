package nodesTelegram

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type TelegramNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newTelegramNode() *TelegramNode {
	return &TelegramNode{
		BaseNode: node.NewBaseNode("telegram", "Telegram"),
	}
}

func (n *TelegramNode) Process(inputs []interface{}) []interface{} {
	message := inputs[0].(string)

	result := "Sent to Telegram: " + message

	return []interface{}{result}
}

func init() {
	node.RegisterProcessor(newTelegramNode())
}
