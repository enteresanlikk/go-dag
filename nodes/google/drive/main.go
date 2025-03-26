package nodesGoogleDrive

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

type GoogleDriveNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

func newGoogleDriveNode() *GoogleDriveNode {
	return &GoogleDriveNode{
		BaseNode: node.NewBaseNode("google-drive", "Google Drive"),
	}
}

func (n *GoogleDriveNode) Process(inputs []interface{}) []interface{} {
	content := inputs[0].(string)

	result := "Uploaded to Google Drive: " + content

	return []interface{}{result}
}

func init() {
	node.RegisterProcessor(newGoogleDriveNode())
}
