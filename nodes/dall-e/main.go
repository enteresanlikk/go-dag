package nodesDallE

import (
	"encoding/base64"
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type DallENode struct {
	node.BaseNode

	Settings interface{}
}

func newDallENode() *DallENode {
	return &DallENode{
		BaseNode: node.NewBaseNode("dall-e", "DALL·E"),
	}
}

func (n *DallENode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("DallENode Process", inputs)

	prompt := inputs["prompt"].(string)

	image := base64.StdEncoding.EncodeToString([]byte("Generated Image for: " + prompt))

	return map[string]interface{}{
		"image": image,
	}
}

func init() {
	node.RegisterProcessor(newDallENode())
}
