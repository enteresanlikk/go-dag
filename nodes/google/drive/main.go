package nodesGoogleDrive

import (
	"github.com/enteresanlikk/go-dag/pkg/node"
)

// type
type GoogleDriveNode struct {
	node.BaseNode

	Settings map[string]interface{}
}

// create
func newGoogleDriveNode() *GoogleDriveNode {
	return &GoogleDriveNode{
		BaseNode: node.NewBaseNode("google-drive", "Google Drive"),
	}
}

// execute
func (n *GoogleDriveNode) Process(inputs []interface{}) []interface{} {
	// DALL-E'den gelen array'in ilk elemanını al
	content := inputs[0].(string)

	// business logic
	result := "Uploaded to Google Drive: " + content

	return []interface{}{result}
}

// init
func init() {
	node.RegisterProcessor(newGoogleDriveNode())
}
