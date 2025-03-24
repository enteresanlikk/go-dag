package nodesCommon

import (
	"time"
)

type NodeOutput struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Inputs        []interface{} `json:"inputs"`
	Outputs       []interface{} `json:"outputs"`
	ExecutionTime time.Duration `json:"executionTime"`
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
	now := time.Now()
	result := execute(inputs)
	elapsed := time.Since(now)

	outputs := []NodeOutput{{
		ID:            n.ID,
		Name:          n.Name,
		Inputs:        inputs,
		Outputs:       result,
		ExecutionTime: elapsed,
	}}

	if len(n.Next) > 0 {
		for _, nextNode := range n.Next {
			nextOutputs := nextNode.Execute(result)
			outputs = append(outputs, nextOutputs...)
		}
	}

	return outputs
}
