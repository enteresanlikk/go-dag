package nodes

import (
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
	_ "github.com/enteresanlikk/go-dag/nodes/dall-e"
	_ "github.com/enteresanlikk/go-dag/nodes/google/drive"
	_ "github.com/enteresanlikk/go-dag/nodes/openai"
	_ "github.com/enteresanlikk/go-dag/nodes/slack"
	_ "github.com/enteresanlikk/go-dag/nodes/telegram"
)

func GetNodeFactory() *nodesCommon.NodeFactory {
	factory := nodesCommon.GetFactory()
	return factory
}
