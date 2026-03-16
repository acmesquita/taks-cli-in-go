# Task tracker

Task tracker is a CLI project to track and manage your tasks using a local JSON file.

---

## How to use (PT-BR)

### Instalação

- **Requisitos**:
  - Go 1.22+ instalado
  - `make` (opcional, mas recomendado)

Clone o repositório e entre na pasta:

```bash
git clone https://github.com/acmesquita/task_tracker.git
cd task_tracker
```

Compile o binário:

```bash
make build
```

O binário `task-cli` será gerado (veja o `Makefile` para detalhes).  
Você pode executar diretamente com `go run` também:

```bash
go run ./cmd/task-cli <command> [flags]
```

### Uso básico

Estrutura geral:

```bash
task-cli <command> [--flags]
```

Comandos disponíveis:

- `add`             – adiciona uma nova tarefa
- `update`          – atualiza a descrição de uma tarefa
- `delete`          – remove uma tarefa
- `list`            – lista tarefas (com ou sem filtro de status)
- `get`             – busca uma tarefa pelo ID
- `mark-done`       – marca uma tarefa como concluída
- `mark-in-progress` – marca uma tarefa como em andamento
- `help`            – exibe a ajuda

### Exemplos de comandos

```bash
# Adicionando uma nova tarefa
task-cli add --description "Buy groceries"

# Atualizando uma tarefa
task-cli update --id 1 --description "Buy groceries and cook dinner"

# Apagando uma tarefa
task-cli delete --id 1

# Marcando tarefa como em andamento ou concluída
task-cli mark-in-progress --id 1
task-cli mark-done --id 1

# Listando todas as tarefas
task-cli list

# Listando tarefas por status
task-cli list --status done
task-cli list --status todo
task-cli list --status in-progress

# Obtendo uma tarefa específica
task-cli get --id 1
```

As tarefas são salvas em um arquivo JSON no diretório atual (veja `core/repository/json_task_repository.go`).

---

## How to use (EN)

### Installation

- **Requirements**:
  - Go 1.22+ installed
  - `make` (optional but recommended)

Clone the repository and enter the folder:

```bash
git clone https://github.com/acmesquita/task_tracker.git
cd task_tracker
```

Build the binary:

```bash
make build
```

The `task-cli` binary will be generated (check the `Makefile` for details).  
You can also run it directly with Go:

```bash
go run ./cmd/task-cli <command> [flags]
```

### Basic usage

General structure:

```bash
task-cli <command> [--flags]
```

Available commands:

- `add`             – add a new task
- `update`          – update a task description
- `delete`          – delete a task
- `list`            – list tasks (optionally filtered by status)
- `get`             – get a task by ID
- `mark-done`       – mark a task as done
- `mark-in-progress` – mark a task as in progress
- `help`            – show help message

### Command examples

```bash
# Adding a new task
task-cli add --description "Buy groceries"

# Updating a task
task-cli update --id 1 --description "Buy groceries and cook dinner"

# Deleting a task
task-cli delete --id 1

# Marking task as in progress or done
task-cli mark-in-progress --id 1
task-cli mark-done --id 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list --status done
task-cli list --status todo
task-cli list --status in-progress

# Getting a specific task
task-cli get --id 1
```

Tasks are stored in a JSON file in the current directory (see `core/repository/json_task_repository.go`).

---

## Requirements
The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:

- [x] Add, Update, and Delete tasks
- [x] Mark a task as in progress or done
- [x] List all tasks
- [x] List all tasks that are done
- [x] List all tasks that are not done
- [x] List all tasks that are in progress

Here are some constraints to guide the implementation:
- [x] You can use any programming language to build this project.
- [x] Use positional arguments in command line to accept user inputs.
- [x] Use a JSON file to store the tasks in the current directory.
- [x] The JSON file should be created if it does not exist.
- [x] Use the native file system module of your programming language to interact with the JSON file.
- [x] Do not use any external libraries or frameworks to build this project.
- [x] Ensure to handle errors and edge cases gracefully.

### Example
The list of commands and their usage is given below:

```bash

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1
# Listing all tasks
task-cli list
# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Task Properties

Each task should have the following properties:

- `id`: A unique identifier for the task
- `description`: A short description of the task
- `status`: The status of the task (todo, in-progress, done)
- `createdAt`: The date and time when the task was created
- `updatedAt`: The date and time when the task was last updated

Make sure to add these properties to the JSON file when adding a new task and update them when updating a task.
