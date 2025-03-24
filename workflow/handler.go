package workflow

import (
	"github.com/enteresanlikk/go-dag/nodes"
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
	"github.com/gofiber/fiber/v2"
)

func ExecuteWorkflowHandler(c *fiber.Ctx) error {
	// get payload from request
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	// get factory
	factory := nodes.GetNodeFactory()

	// create nodes
	nodes := make(map[string]nodesCommon.Node)
	for _, payloadNode := range payload.Nodes {
		node, err := factory.Create(payloadNode.ID, payloadNode.Settings)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating node",
			})
		}
		nodes[payloadNode.ID] = node
	}

	// set edges
	for _, edge := range payload.Edges {
		if nodes[edge.Source] == nil || nodes[edge.Target] == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Node not found",
			})
		}

		sourceNode := nodes[edge.Source]
		targetNode := nodes[edge.Target]

		sourceNode.SetChild(targetNode)
		targetNode.SetParent(sourceNode)
	}

	// get start node
	rootPayloadNode := payload.Nodes[0]
	rootNode := nodes[rootPayloadNode.ID]

	// execute workflow
	result := rootNode.Execute(rootPayloadNode.Data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Workflow executed successfully",
		"result":  result,
	})
}
