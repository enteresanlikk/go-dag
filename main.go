package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enteresanlikk/go-dag/workflow"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Post("/workflow", workflow.ExecuteWorkflowHandler)

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
