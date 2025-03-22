package nodesCommon

import (
	"fmt"
)

type NodeOutput struct {
	NodeName string
	Data     interface{}
}

type Node interface {
	Execute(inputs ...interface{}) []NodeOutput
	SetNext(node Node)
	GetName() string
	GetNext() []Node
}

type BaseNode struct {
	Enabled bool
	Next    []Node
	Name    string
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

func (n *BaseNode) ExecuteWithCheck(inputs []interface{}, execute func([]interface{}) interface{}) []NodeOutput {
	fmt.Printf("Executing node: %s\n", n.Name)

	if !n.Enabled {
		fmt.Printf("%s Node is disabled, skipping...\n", n.Name)
		return []NodeOutput{{NodeName: n.Name, Data: inputs[0]}}
	}

	result := execute(inputs)
	outputs := []NodeOutput{{NodeName: n.Name, Data: result}}

	if len(n.Next) > 0 {
		for _, nextNode := range n.Next {
			nextOutputs := nextNode.Execute(result)
			outputs = append(outputs, nextOutputs...)
		}
	}

	return outputs
}
