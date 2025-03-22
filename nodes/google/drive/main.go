package nodesGoogleDrive

import (
	"fmt"
	"time"

	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
)

type GoogleDriveNode struct {
	nodesCommon.BaseNode

	Folder string
}

func NewGoogleDriveNode(folder string) *GoogleDriveNode {
	return &GoogleDriveNode{
		BaseNode: nodesCommon.BaseNode{
			Enabled: true,
			Name:    "Google Drive",
		},
		Folder: folder,
	}
}

func (n *GoogleDriveNode) Execute(inputs ...interface{}) []nodesCommon.NodeOutput {
	return n.ExecuteWithCheck(inputs, func(inputs []interface{}) interface{} {
		time.Sleep(2 * time.Second)

		imageData := inputs[0].(string)
		savedPath := fmt.Sprintf("%s/image_%s.png", n.Folder, imageData[:10])
		fmt.Println("Image saved to Google Drive:", savedPath)

		return savedPath
	})
}
