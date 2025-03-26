package nodesGoogleDrive

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/pkg/node"
)

type GoogleDriveNode struct {
	node.BaseNode
}

func NewGoogleDriveNode() *GoogleDriveNode {
	return &GoogleDriveNode{
		BaseNode: node.NewBaseNode("google-drive", "Google Drive"),
	}
}

func (n *GoogleDriveNode) Process(inputs map[string]interface{}) map[string]interface{} {
	fmt.Println("GoogleDriveNode Process", inputs)

	file := inputs["file"].(string)

	result := "Uploaded to Google Drive: " + file

	return map[string]interface{}{
		"result": result,
	}
}

func init() {
	node.RegisterProcessor(NewGoogleDriveNode())
}
