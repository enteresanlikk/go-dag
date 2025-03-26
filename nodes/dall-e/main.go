package nodesDallE

import (
	"encoding/base64"
	"time"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type DallENode struct {
	node.BaseNode

	Settings interface{}
}

func newDallENode() *DallENode {
	return &DallENode{
		BaseNode: node.NewBaseNode("dall-e", "DALL·E"),
	}
}

// execute
func (n *DallENode) Process(inputs []interface{}) []interface{} {
	time.Sleep(10 * time.Second)

	// OpenAI'dan gelen array'in ilk elemanını al
	prompt := inputs[0].(string)

	//business logic
	image := base64.StdEncoding.EncodeToString([]byte("Generated Image for: " + prompt))

	return []interface{}{image}
}

// init
func init() {
	node.RegisterProcessor(newDallENode())
}
