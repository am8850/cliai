# cliai - CLI AI Commander

A utility written in Go to take natural language commands, generate git, Azure CLI and kubernetes commands, and execute them using models in Azure, OpenAI, and locally with Ollama.

## Configuration

Create a file `` with the following settings:
```json
{
    "endpoint": "https://<NAME>.openai.azure.com/openai/deployments/gpt-4o/chat/completions?api-version=2025-01-01-preview",
    "key": "<KEY>",
    "model": "gpt-4o",
    "type": "azure",
    "temperature": 0.1
}
```
Where:
- Type: azure or openai

> Note: for Ollama, set the endpoint, and for the key enter `123`

## Easy install - WSL, Ubuntu

- For easy install type: `make install`

## Snapshot

![Snapshot of cliai showing the output from executing git commands](images/cliai-snapshot-02.png)

## Usage

- Git CLI (Git compound command):
  - `cliai git -p "What is the git version and list all branches"`

- Docker CLI
  - `cliai docker -p "List all running containers"`

- Azure CLI
  - `cliai az -p "Get the first resource group in eastus as a table"`

- kubectl CLI
  - `cliai k -p "Get the all pods in all namespaces"`

- Scaffold
  - `cliai sc -p "Generate a Python FAST API to manage customer"`

- Refactor
  - `cliai re -f app.py -o app_new.app`

- Any
  - `cliai any -p "What is the speed of light?"`