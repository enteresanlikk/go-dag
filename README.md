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
            "inputs": {
                "prompt": "Create a futuristic city illustration"
            },
            "settings": {
                "apiKey": "OPENAI_API_KEY"
            }
        },
        {
            "id": "dall-e",
            "inputs": {
                "prompt": "$openai.response"
            },
            "settings": {
                "apiKey": "DALL_E_API_KEY"
            }
        },
        {
            "id": "google-drive",
            "inputs": {
                "file": "bu nasıl bir file $dall-e.image"
            },
            "settings": {
                "folder": "GOOGLE_DRIVE_FOLDER",
                "apiKey": "GOOGLE_DRIVE_API_KEY"
            }
        },
        {
            "id": "slack",
            "inputs": {
                "message": "slack Image created: $google-drive.result"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "telegram Image created: $google-drive.result"
            },
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
}
```

```json
{
    "nodes": [
        {
            "id": "condition",
            "inputs": {
                "value": "test-value"
            },
            "settings": {
                "condition_type": "equals",
                "expected_value": "test-value"
            }
        },
        {
            "id": "slack",
            "inputs": {
                "message": "$condition.true_value"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "$condition.false_value"
            },
            "settings": {
                "botToken": "TELEGRAM_BOT_TOKEN",
                "chatId": "TELEGRAM_CHAT_ID"
            }
        }
    ],
    "edges": [
        {
            "source": "condition",
            "target": "slack"
        },
        {
            "source": "condition",
            "target": "telegram"
        }
    ]
}
```

```json
{
    "nodes": [
        {
            "id": "openai",
            "inputs": {
                "prompt": "Create a futuristic city illustration"
            },
            "settings": {
                "apiKey": "OPENAI_API_KEY"
            }
        },
        {
            "id": "dall-e",
            "inputs": {
                "prompt": "$openai.response"
            },
            "settings": {
                "apiKey": "DALL_E_API_KEY"
            }
        },
        {
            "id": "google-drive",
            "inputs": {
                "file": "bu nasıl bir file $dall-e.image"
            },
            "settings": {
                "folder": "GOOGLE_DRIVE_FOLDER",
                "apiKey": "GOOGLE_DRIVE_API_KEY"
            }
        },
        {
            "id": "slack",
            "inputs": {
                "message": "slack Image created: $google-drive.result"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "telegram Image created: $google-drive.result"
            },
            "settings": {
                "botToken": "TELEGRAM_BOT_TOKEN",
                "chatId": "TELEGRAM_CHAT_ID"
            }
        },
        {
            "id": "merge",
            "inputs": {
                "telegram": "$telegram.message",
                "slack": "$slack.message",
                "dall-e": "$dall-e.image",
                "otherInput": "other input value"
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
