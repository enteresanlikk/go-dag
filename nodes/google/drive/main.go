package nodesGoogleDrive

import (
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

// type
type GoogleDriveNode struct {
	nodesCommon.BaseNode

	Settings map[string]interface{}
}

// base node settings
var baseNode = nodesCommon.BaseNode{
	ID:   "google-drive",
	Name: "Google Drive",
}

// create
func NewGoogleDriveNode(settings map[string]interface{}) *GoogleDriveNode {
	return &GoogleDriveNode{
		BaseNode: baseNode,
		Settings: settings,
	}
}

// execute
func (n *GoogleDriveNode) Execute(inputs []interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) []interface{} {
		// time.Sleep(2 * time.Second)

		imageData := inputs[0].(string)
		folder := n.Settings["folder"].(string)

		//business logic
		savedPath := folder + "/image_" + imageData + ".png"

		return []interface{}{savedPath}
	})
}

// factory
func CreateGoogleDriveNode(settings map[string]interface{}) (nodesCommon.Node, error) {
	return NewGoogleDriveNode(settings), nil
}

// init
func init() {
	factory := nodesCommon.GetFactory()
	factory.Register(baseNode.ID, CreateGoogleDriveNode)
}
