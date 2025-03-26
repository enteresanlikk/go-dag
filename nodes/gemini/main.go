package nodesGemini

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type GeminiNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newGeminiNode() *GeminiNode {
	return &GeminiNode{
		BaseNode: node.NewBaseNode("gemini", "Gemini"),
	}
}

func (n *GeminiNode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("GeminiNode Process", inputs)

	prompt := inputs["prompt"].(string)

	return map[string]interface{}{
		"prompt": prompt,
	}
}

func init() {
	node.RegisterProcessor(newGeminiNode())
}
