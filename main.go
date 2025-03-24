package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enteresanlikk/go-dag/nodes"
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	ID       string                 `json:"id"`
	Data     []interface{}          `json:"data"`
	Settings map[string]interface{} `json:"settings"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// handler
func executeWorkflowHandler(c *fiber.Ctx) error {
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

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Post("/workflow", executeWorkflowHandler)

	go func() {
		log.Printf("Server starting on http://%s:%s", host, port)
		if err := app.Listen(host + ":" + port); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	_, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
