package nodesTelegram

import (
	"time"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type TelegramNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

// create
func newTelegramNode() *TelegramNode {
	return &TelegramNode{
		BaseNode: node.NewBaseNode("telegram", "Telegram"),
	}
}

// execute
func (n *TelegramNode) Process(inputs []interface{}) []interface{} {
	time.Sleep(1 * time.Second)

	// Google Drive'dan gelen array'in ilk elemanını al
	message := inputs[0].(string)

	// business logic
	result := "Sent to Telegram: " + message

	return []interface{}{result}
}

// init
func init() {
	node.RegisterProcessor(newTelegramNode())
}
