package nodesOpenAI

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type MergeNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newMergeNode() *MergeNode {
	return &MergeNode{
		BaseNode: node.NewBaseNode("merge", "Merge"),
	}
}

func (n *MergeNode) Process(inputs []interface{}) []interface{} {
	for index, input := range inputs {
		fmt.Println("MergeNode Process", index, input)
	}

	return inputs
}

func init() {
	node.RegisterProcessor(newMergeNode())
}
