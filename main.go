package main

import (
	"fmt"

	nodesDallE "github.com/enteresanlikk/go-dag/nodes/dall-e"
	nodesGoogleDrive "github.com/enteresanlikk/go-dag/nodes/google/drive"
	nodesOpenAI "github.com/enteresanlikk/go-dag/nodes/openai"
	nodesSlack "github.com/enteresanlikk/go-dag/nodes/slack"
)

func main() {
	openAI := nodesOpenAI.NewOpenAINode("OPEN_AI_API_KEY")

	dalle := nodesDallE.NewDallENode("DALL_E_API_KEY")

	googleDrive := nodesGoogleDrive.NewGoogleDriveNode("Google Drive Folder")

	slack := nodesSlack.NewSlackNode("SLACK_WEBHOOK_URL")

	openAI.SetNext(dalle)
	dalle.SetNext(googleDrive)
	googleDrive.SetNext(slack)

	inputPrompt := "Create a futuristic city illustration"
	fmt.Println("🚀 Starting DAG Workflow...")
	finalOutput := openAI.Execute(inputPrompt)

	fmt.Println("\n✅ Workflow Completed! Final Output:", finalOutput)
}
