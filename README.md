# Go Directed Acyclic Graph (DAG) [Wikipedia](https://en.wikipedia.org/wiki/Directed_acyclic_graph)

## Run

```bash
docker compose up -d
```

## Stop

```bash
docker compose down
```

## Workflow

```json
{
    "nodes": [
        {
            "id": "openai",
            "inputs": [
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
        },
        {
            "id": "merge"
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
        },
        {
            "source": "slack",
            "target": "merge"
        },
        {
            "source": "telegram",
            "target": "merge"
        }
    ]
}
```

## Run Workflow

```bash
curl --location 'http://127.0.0.1:3000/workflow' --header 'Content-Type: application/json' --data @workflow.json
```
