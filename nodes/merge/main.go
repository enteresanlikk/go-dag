package nodesOpenAI

import (
	"fmt"
	"time"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type MergeNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

// create
func newMergeNode() *MergeNode {
	return &MergeNode{
		BaseNode: node.NewBaseNode("merge", "Merge"),
	}
}

// execute
func (n *MergeNode) Process(inputs []interface{}) []interface{} {
	time.Sleep(5 * time.Second)

	for _, input := range inputs {
		fmt.Println("bilal", input)
	}

	return inputs
}

// init
func init() {
	node.RegisterProcessor(newMergeNode())
}
