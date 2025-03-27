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
                "prompt": "$id[openai].outputs[response]"
            },
            "settings": {
                "apiKey": "DALL_E_API_KEY"
            }
        },
        {
            "id": "google-drive",
            "inputs": {
                "file": "bu nasıl bir file $id[dall-e].outputs[image]"
            },
            "settings": {
                "folder": "GOOGLE_DRIVE_FOLDER",
                "apiKey": "GOOGLE_DRIVE_API_KEY"
            }
        },
        {
            "id": "slack",
            "inputs": {
                "message": "slack Image created: $id[google-drive].outputs[result]"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "telegram Image created: $id[google-drive].outputs[result]"
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
                "message": "$id[condition].outputs[true_value]"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "$id[condition].outputs[false_value]"
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
                "prompt": "$id[openai].outputs[response]"
            },
            "settings": {
                "apiKey": "DALL_E_API_KEY"
            }
        },
        {
            "id": "google-drive",
            "inputs": {
                "file": "bu nasıl bir file $id[dall-e].outputs[image]"
            },
            "settings": {
                "folder": "GOOGLE_DRIVE_FOLDER",
                "apiKey": "GOOGLE_DRIVE_API_KEY"
            }
        },
        {
            "id": "slack",
            "inputs": {
                "message": "slack Image created: $id[google-drive].outputs[result]"
            },
            "settings": {
                "webhook": "SLACK_WEBHOOK"
            }
        },
        {
            "id": "telegram",
            "inputs": {
                "message": "telegram Image created: $id[google-drive].outputs[result]"
            },
            "settings": {
                "botToken": "TELEGRAM_BOT_TOKEN",
                "chatId": "TELEGRAM_CHAT_ID"
            }
        },
        {
            "id": "merge",
            "inputs": {
                "telegram": "$id[telegram].outputs[message]",
                "slack": "$id[slack].outputs[message]",
                "dall-e": "$id[dall-e].outputs[image]",
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
