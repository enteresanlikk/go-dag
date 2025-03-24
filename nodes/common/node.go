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
	SetChild(node Node)
	GetChildren() []Node
	SetParent(node Node)
	GetParents() []Node
}

type BaseNode struct {
	ID       string
	Name     string
	Children []Node
	Parents  []Node
}

func (n *BaseNode) SetChild(node Node) {
	n.Children = append(n.Children, node)
}

func (n *BaseNode) GetChildren() []Node {
	return n.Children
}

func (n *BaseNode) SetParent(node Node) {
	n.Parents = append(n.Parents, node)
}

func (n *BaseNode) GetParents() []Node {
	return n.Parents
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

	if len(n.Children) > 0 {
		for _, childNode := range n.Children {
			childOutputs := childNode.Execute(result)
			outputs = append(outputs, childOutputs...)
		}
	}

	return outputs
}
