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

func (n *DallENode) Execute(input interface{}) interface{} {
	return n.ExecuteWithCheck(input, func(input interface{}) interface{} {
		time.Sleep(10 * time.Second)

		prompt := input.(string)
		image := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Generated Image for: %s", prompt)))

		fmt.Println("DALL·E Output (Base64):", image[:30]+"...")

		if n.Next != nil {
			return n.Next.Execute(image)
		}
		return image
	})
}
