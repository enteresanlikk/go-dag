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

func (n *BaseNode) ExecuteWithCheck(inputs []interface{}, execute func([]interface{}) interface{}) []NodeOutput {
	name := n.GetName()

	fmt.Printf("Executing node: %s\n", name)

	if !n.Enabled {
		fmt.Printf("%s Node is disabled, skipping...\n", name)
		return []NodeOutput{{NodeName: name, Data: inputs[0]}}
	}

	result := execute(inputs)
	outputs := []NodeOutput{{NodeName: name, Data: result}}

	if len(n.Next) > 0 {
		for _, nextNode := range n.Next {
			nextOutputs := nextNode.Execute(result)
			outputs = append(outputs, nextOutputs...)
		}
	}

	return outputs
}
