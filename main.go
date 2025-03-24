package main

import (
	"fmt"

	"github.com/enteresanlikk/go-dag/nodes"
	nodesCommon "github.com/enteresanlikk/go-dag/nodes/common"
	"github.com/goccy/go-json"
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

func GetPayload() (Payload, error) {
	payloadString := `{
		"nodes": [
			{
				"id": "openai",
				"data": [
					"Create a futuristic city illustration"
				],
				"settings": {
					"apiKey": "OPENAI_API_KEY"
				}
			},
			{
				"id": "dall-e",
				"settings": {
					"apiKey": "DALL_E_API_KEY"
				}
			},
			{
				"id": "google-drive",
				"settings": {
					"folder": "GOOGLE_DRIVE_FOLDER",
					"apiKey": "GOOGLE_DRIVE_API_KEY"
				}
			},
			{
				"id": "slack",
				"settings": {
					"webhook": "SLACK_WEBHOOK"
				}
			},
			{
				"id": "telegram",
				"settings": {
					"botToken": "TELEGRAM_BOT_TOKEN",
					"chatId": "TELEGRAM_CHAT_ID"
				}
			}
		],
		"edges": [
			{
				"source": "openai",
				"target": "dall-e"
			},
			{
				"source": "dall-e",
				"target": "google-drive"
			},
			{
				"source": "google-drive",
				"target": "slack"
			},
			{
				"source": "google-drive",
				"target": "telegram"
			}
		]
	}`

	var payload Payload
	err := json.Unmarshal([]byte(payloadString), &payload)
	if err != nil {
		return Payload{}, err
	}

	return payload, nil
}

func main() {
	// get payload
	payload, err := GetPayload()
	if err != nil {
		fmt.Println("Error getting payload:", err)
		return
	}

	// get factory
	factory := nodes.GetNodeFactory()

	// create nodes
	nodes := make(map[string]nodesCommon.Node)
	for _, payloadNode := range payload.Nodes {
		node, err := factory.Create(payloadNode.ID, payloadNode.Settings)
		if err != nil {
			fmt.Println("Error creating node:", err)
			return
		}
		nodes[payloadNode.ID] = node
	}

	// set edges
	for _, edge := range payload.Edges {
		if nodes[edge.Source] == nil || nodes[edge.Target] == nil {
			fmt.Println("Node not found:", edge.Source, edge.Target)
			return
		}
		nodes[edge.Source].SetNext(nodes[edge.Target])
	}

	// execute workflow
	nodes[payload.Nodes[0].ID].Execute(payload.Nodes[0].Data)
}
