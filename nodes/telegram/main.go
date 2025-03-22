package nodesTelegram

import (
	"fmt"
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

type TelegramNode struct {
	nodesCommon.BaseNode

	BotToken string
	ChatID   string
}

func NewTelegramNode(botToken string, chatID string) *TelegramNode {
	return &TelegramNode{
		BaseNode: nodesCommon.BaseNode{
			Enabled: true,
			Name:    "Telegram",
		},
		BotToken: botToken,
		ChatID:   chatID,
	}
}

func (n *TelegramNode) Execute(inputs ...interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) interface{} {
		time.Sleep(1 * time.Second)

		message := fmt.Sprintf("New AI-generated image saved at: %s", inputs[0].(string))
		fmt.Printf("Sending Telegram Message to chat %s: %s\n", n.ChatID, message)

		return message
	})
}
