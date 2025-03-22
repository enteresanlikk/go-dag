package nodesDallE

import (
	"encoding/base64"
	"fmt"
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

type DallENode struct {
	nodesCommon.BaseNode

	ApiKey string
}

func NewDallENode(apiKey string) *DallENode {
	return &DallENode{
		BaseNode: nodesCommon.BaseNode{
			Enabled: true,
			Name:    "DALL·E",
		},
		ApiKey: apiKey,
	}
}

func (n *DallENode) Execute(inputs ...interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) interface{} {
		time.Sleep(10 * time.Second)

		prompt := inputs[0].(string)
		image := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Generated Image for: %s", prompt)))

		fmt.Println("DALL·E Output (Base64):", image[:30]+"...")

		return image
	})
}
