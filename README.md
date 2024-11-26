# cliai - CLI Commander

A utility written in Go to take natural language commands, generate git, Azure CLI and kubernetes commands, and execute them.

## Snapshot

![Snapshot of cliai showing the output from executing git commands](images/cliai-snapshot-02.png)

## Usage

- Git CLI (Git compound command):
  - `cliai git -p "What is the git version and list all branches"`

- Azure CLI
  - `cliai az -p "Get the first resource group in eastus as a table"`

- kubectl CLI
  - `cliai k8s -p "Get the all pods in all namespaces"`
