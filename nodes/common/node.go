package nodesCommon

import (
	"fmt"
	"time"
)

type NodeOutput struct {
	NodeName string
	Data     []interface{}
}

type Node interface {
	Execute(inputs []interface{}) []NodeOutput
	SetNext(node Node)
	GetName() string
	GetNext() []Node
}

type BaseNode struct {
	ID   string
	Name string
	Next []Node
}

func (n *BaseNode) SetNext(node Node) {
	n.Next = append(n.Next, node)
}

func (n *BaseNode) GetName() string {
	return n.Name
}

func (n *BaseNode) GetNext() []Node {
	return n.Next
}

func (n *BaseNode) ExecuteWithCheck(inputs []interface{}, execute func([]interface{}) []interface{}) []NodeOutput {
	name := n.GetName()

	fmt.Printf("Executing node: %s\n", name)

	now := time.Now()
	result := execute(inputs)
	elapsed := time.Since(now)

	fmt.Printf("Executed node: %s\n", name)
	fmt.Printf("Execution time: %s\n", elapsed)
	fmt.Println("--------------------------------")

	outputs := []NodeOutput{{NodeName: name, Data: result}}

	if len(n.Next) > 0 {
		for _, nextNode := range n.Next {
			nextOutputs := nextNode.Execute(result)
			outputs = append(outputs, nextOutputs...)
		}
	}

	return outputs
}
