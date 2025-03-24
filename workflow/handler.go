package workflow

import (
	"fmt"

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
			fmt.Println("Error creating node:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating node",
			})
		}
		nodes[payloadNode.ID] = node
	}

	// set edges
	for _, edge := range payload.Edges {
		if nodes[edge.Source] == nil || nodes[edge.Target] == nil {
			fmt.Println("Node not found:", edge.Source, edge.Target)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Node not found",
			})
		}
		nodes[edge.Source].SetNext(nodes[edge.Target])
	}

	// execute workflow
	result := nodes[payload.Nodes[0].ID].Execute(payload.Nodes[0].Data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Workflow executed successfully",
		"result":  result,
	})
}
