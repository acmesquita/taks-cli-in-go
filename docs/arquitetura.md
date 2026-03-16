## Arquitetura

A aplicação segue uma arquitetura em camadas, separando claramente parsing de comandos, orquestração e regras de negócio.

- **Interface de Linha de Comando (CLI / Commands)**
  - Arquivo principal: `infra/commands/commands.go`
  - Responsável por:
    - Validar o comando recebido (`add`, `update`, `delete`, `list`, `get`, `mark-done`, `mark-in-progress`, `help`).
    - Exibir a mensagem de ajuda (`HandleHelperMessage`).
    - Interpretar argumentos e opções da linha de comando, convertendo-os em um `map[string]string` via `ParseOptions`.

- **Processador de Comandos (Application Layer)**
  - Arquivo: `infra/processor/processor.go`
  - Estrutura central: `Processor`, que recebe um `services.TaskService`.
  - Função `Process`:
    - Constrói um `adapter.Request` a partir do comando e opções.
    - Faz o `switch` do comando e delega a chamada para os *handlers* específicos (`AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `FindTask`, `MarkDoneTask`, `MarkInProgressTask`).
    - Trata o comando `help` chamando `commands.HandleHelperMessage`.

- **Adaptador de Request**
  - Arquivo: `infra/processor/adapter/request.go`
  - Tipo `Request` encapsula:
    - `Command` (string)
    - `Options` (map de opções já parseadas)
  - Provê métodos de acesso (`GetCommand`, `GetOptions`) para desacoplar o processador dos detalhes de como as opções foram montadas.

- **Camada de Handlers (Interface com a Aplicação)**
  - Pasta: `infra/processor/handlers`
  - Cada operação do domínio tem um handler dedicado:
    - `AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `FindTask`, `MarkDoneTask`, `MarkInProgressTask`.
  - Responsabilidades dos handlers:
    - Validar a presença dos parâmetros obrigatórios (por exemplo, `id`, `description`).
    - Chamar os métodos apropriados do `TaskService`.
    - Tratar cenários de erro (como tarefa não encontrada) imprimindo mensagens amigáveis e definindo o código de saída se necessário.
    - Imprimir o resultado das operações (ID, descrição, status e mensagens de sucesso).
    - No caso do comando `list`, aceitar um parâmetro opcional `--status` (`todo`, `in-progress`, `done` ou `all`) para filtrar as tarefas por status, usando `service.ListTasksByStatus`.

- **Camada de Domínio / Serviços (Core)**
  - Referenciada como `github.com/acmesquita/task_tracker/core/services`.
  - A interface `TaskService` concentra a lógica de negócio de tarefas:
    - `AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `GetTask`, `MarkTaskAsDone`, `MarkTaskAsInProgress`, `ListTasksByStatus`.
  - A infra (CLI/handlers) não conhece os detalhes de persistência (arquivo, memória, DB etc.), apenas interage com o serviço.

- **Persistência (Repositório de Tarefas)**
  - Implementação padrão: `core/repository/json_task_repository.go` (`JSONTaskRepository`).
  - Responsável por carregar e salvar as tarefas em um arquivo JSON:
    - Cria o arquivo caso não exista.
    - Exponde operações de repositório (`AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `GetTask`) usadas pelo `TaskService`.

---

## Architecture (EN)

The application uses a layered architecture with a clear separation between command parsing, orchestration, and business logic.

- **Command-Line Interface (CLI / Commands)**
  - Main file: `infra/commands/commands.go`
  - Responsibilities:
    - Validate the incoming command (`add`, `update`, `delete`, `list`, `get`, `mark-done`, `mark-in-progress`, `help`).
    - Display the help message (`HandleHelperMessage`).
    - Parse CLI arguments and options into a `map[string]string` via `ParseOptions`.

- **Command Processor (Application Layer)**
  - File: `infra/processor/processor.go`
  - Core struct: `Processor`, which receives a `services.TaskService`.
  - `Process` function:
    - Builds an `adapter.Request` from command and options.
    - Uses a `switch` on the command and dispatches to the proper handler (`AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `FindTask`, `MarkDoneTask`, `MarkInProgressTask`).
    - Handles the `help` command by delegating to `commands.HandleHelperMessage`.

- **Request Adapter**
  - File: `infra/processor/adapter/request.go`
  - `Request` type encapsulates:
    - `Command` (string)
    - `Options` (parsed options map)
  - Provides accessor methods (`GetCommand`, `GetOptions`) to decouple the processor from how options are built.

- **Handlers Layer (Application Interface)**
  - Folder: `infra/processor/handlers`
  - Each domain operation has a dedicated handler:
    - `AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `FindTask`, `MarkDoneTask`, `MarkInProgressTask`.
  - Handler responsibilities:
    - Validate required parameters (`id`, `description`, etc.).
    - Call the proper `TaskService` method.
    - Handle error scenarios (e.g., task not found) with user-friendly messages and proper exit codes.
    - Print operation results (ID, description, status, and success messages).
    - For the `list` command, accept an optional `--status` argument (`todo`, `in-progress`, `done`, or `all`) to filter tasks by status via `service.ListTasksByStatus`.

- **Domain / Services Layer (Core)**
  - Referenced as `github.com/acmesquita/task_tracker/core/services`.
  - `TaskService` interface encapsulates the task business logic:
    - `AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `GetTask`, `MarkTaskAsDone`, `MarkTaskAsInProgress`, `ListTasksByStatus`.
  - The infra (CLI/handlers) does not know about persistence details (file, memory, DB, etc.), it only talks to the service.

- **Persistence (Task Repository)**
  - Default implementation: `core/repository/json_task_repository.go` (`JSONTaskRepository`).
  - Responsible for loading and saving tasks to a JSON file:
    - Creates the file when it does not exist.
    - Exposes repository operations (`AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `GetTask`) used by `TaskService`.

