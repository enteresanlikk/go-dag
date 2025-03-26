package nodesTelegram

import (
	"fmt"

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

func (n *TelegramNode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("TelegramNode Process", inputs)

	message := inputs["message"].(string)

	return map[string]interface{}{
		"message": message,
	}
}

func init() {
	node.RegisterProcessor(newTelegramNode())
}
