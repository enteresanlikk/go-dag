package nodesCondition

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type ConditionNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newConditionNode() *ConditionNode {
	return &ConditionNode{
		BaseNode: node.NewBaseNode("condition", "Condition"),
	}
}

func (n *ConditionNode) Process(inputs []interface{}) []interface{} {
	if len(inputs) == 0 {
		return []interface{}{nil, nil}
	}

	conditionType := n.GetSetting("condition_type", "equals").(string)
	expectedValue := n.GetSetting("expected_value", "")

	input := inputs[0]
	var result bool

	switch conditionType {
	case "equals":
		result = input == expectedValue
	case "not_equals":
		result = input != expectedValue
	case "contains":
		if str, ok := input.(string); ok {
			if expStr, ok := expectedValue.(string); ok {
				result = len(str) > 0 && len(expStr) > 0 && str != "" && expStr != "" && str != expStr
			}
		}
	case "greater_than":
		if num, ok := input.(float64); ok {
			if expNum, ok := expectedValue.(float64); ok {
				result = num > expNum
			}
		}
	case "less_than":
		if num, ok := input.(float64); ok {
			if expNum, ok := expectedValue.(float64); ok {
				result = num < expNum
			}
		}
	default:
		result = false
	}

	if result {
		return []interface{}{input, nil}
	}

	return []interface{}{nil, input}
}

func init() {
	node.RegisterProcessor(newConditionNode())
}
