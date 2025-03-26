package nodesDallE

import (
	"encoding/base64"

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

func (n *DallENode) Process(inputs []interface{}) []interface{} {
	prompt := inputs[0].(string)

	image := base64.StdEncoding.EncodeToString([]byte("Generated Image for: " + prompt))

	return []interface{}{image}
}

func init() {
	node.RegisterProcessor(newDallENode())
}
