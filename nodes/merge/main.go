package nodesOpenAI

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type MergeNode struct {
	node.BaseNode
}

func NewMergeNode() *MergeNode {
	return &MergeNode{
		BaseNode: node.NewBaseNode("merge", "Merge"),
	}
}

func (n *MergeNode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("MergeNode Process", inputs)

	return inputs
}

func init() {
	node.RegisterProcessor(NewMergeNode())
}
