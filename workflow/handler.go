package workflow

import (
	"log"

	"github.com/enteresanlikk/go-dag/pkg/graph"
	"github.com/gofiber/fiber/v2"

	_ "github.com/enteresanlikk/go-dag/nodes"
)

func ExecuteWorkflowHandler(c *fiber.Ctx) error {
	// get payload from request
	var payload *graph.GraphConfig
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	g := graph.NewGraph()

	if err := g.LoadFromJSON(payload); err != nil {
		log.Fatalf("Failed to load graph from JSON: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Workflow executed successfully",
	})
}
