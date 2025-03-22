package nodesCommon

import (
	"fmt"
)

type Node interface {
	Execute(input interface{}) interface{}
	SetNext(node Node)
	GetName() string
}

type BaseNode struct {
	Enabled bool
	Next    Node
	Name    string
}

func (n *BaseNode) SetNext(node Node) {
	n.Next = node
}

func (n *BaseNode) GetName() string {
	return n.Name
}

func (n *BaseNode) ExecuteWithCheck(input interface{}, execute func(interface{}) interface{}) interface{} {
	fmt.Println("Executing node:", n.Name)

	if !n.Enabled {
		fmt.Printf("%s Node is disabled, skipping...\n", n.Name)
		return input
	}

	return execute(input)
}
