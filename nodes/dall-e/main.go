package nodesDallE

import (
	"encoding/base64"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

// type
type DallENode struct {
	nodesCommon.BaseNode

	Settings interface{}
}

// base node settings
var baseNode = nodesCommon.BaseNode{
	ID:   "dall-e",
	Name: "DALL·E",
}

// create
func NewDallENode(settings map[string]interface{}) *DallENode {
	return &DallENode{
		BaseNode: baseNode,
		Settings: settings,
	}
}

// execute
func (n *DallENode) Execute(inputs []interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) []interface{} {
		// time.Sleep(10 * time.Second)

		prompt := inputs[0].(string)

		//business logic
		image := base64.StdEncoding.EncodeToString([]byte("Generated Image for: " + prompt))

		return []interface{}{image}
	})
}

// factory
func CreateDallENode(settings map[string]interface{}) (nodesCommon.Node, error) {
	return NewDallENode(settings), nil
}

// init
func init() {
	factory := nodesCommon.GetFactory()
	factory.Register(baseNode.ID, CreateDallENode)
}
