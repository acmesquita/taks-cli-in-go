## Contexto para AIs / Skills (Task Tracker)

### Objetivo do projeto

PT: CLI para gerenciar tarefas (`task-cli`), com comandos como `add`, `update`, `delete`, `list`, `get`, `mark-done`, seguindo arquitetura em camadas (CLI, processor, handlers, serviços de domínio).  
EN: CLI task manager (`task-cli`) with commands like `add`, `update`, `delete`, `list`, `get`, `mark-done`, built with a layered architecture (CLI, processor, handlers, domain services).

---

## Arquitetura (resumo rápido)

- **CLI / Commands**
  - Arquivo principal: `infra/commands/commands.go`
  - Responsável por:
    - Validar o comando recebido (`add`, `update`, `delete`, `list`, `get`, `mark-done`, `help`).
    - Exibir a mensagem de ajuda (`HandleHelperMessage`).
    - Interpretar argumentos e opções (`ParseOptions`) em `map[string]string`.

- **Processador de Comandos (Application Layer)**
  - Arquivo: `infra/processor/processor.go`
  - Struct central: `Processor` (recebe `services.TaskService`).
  - Função `Process`:
    - Monta `adapter.Request` com comando + opções.
    - Faz `switch` do comando e delega para handlers (`AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `FindTask`, `MarkDoneTask`).
    - Trata `help` chamando `commands.HandleHelperMessage`.

- **Adaptador de Request**
  - Arquivo: `infra/processor/adapter/request.go`
  - Tipo `Request`:
    - Campos: `Command` (string), `Options` (map).
    - Métodos: `GetCommand`, `GetOptions` para desacoplar o `Processor` de detalhes de parsing.

- **Handlers (Interface com a aplicação)**
  - Pasta: `infra/processor/handlers`
  - Um handler por operação de domínio:
    - `add_task.go`, `update_task.go`, `delete_task.go`, `list_task.go`,
      `find_task.go`, `mark_done_task.go`.
  - Responsabilidades:
    - Validar parâmetros obrigatórios (`id`, `description` etc.).
    - Chamar o método correto de `TaskService`.
    - Tratar erros (ex.: tarefa não encontrada, parâmetros faltando).
    - Imprimir resultados e mensagens de sucesso/erro.

- **Domínio / Serviços (Core)**
  - Pacote referenciado: `github.com/acmesquita/task_tracker/core/services`
  - Interface `TaskService`:
    - `AddTask`, `UpdateTask`, `DeleteTask`, `ListTasks`, `GetTask`, `MarkTaskAsDone`.
  - A infra não sabe sobre persistência; só conversa com o serviço.

---

## Decisões de solução importantes

- **Validação antecipada de comandos e opções**
  - `ParseCommands` / `validateCommands` → só comandos suportados são executados.
  - `ParseOptions` converte args em mapa, suporta `--description` e `-d`.
  - Para `add`/`update`, valida que `description` não é vazia.

- **Tratamento consistente de erros de entrada**
  - Handlers (`AddTask`, `UpdateTask`, `FindTask`, `DeleteTask`, `MarkDoneTask`):
    - Ao faltar parâmetro obrigatório: imprime mensagem clara, chama `HandleHelperMessage`, faz `os.Exit(1)`.

- **Encapsulamento da lógica de negócio**
  - Toda a lógica de tarefas está em `TaskService`.
  - CLI/handlers focam em:
    - Validar entrada
    - Orquestrar chamadas
    - Imprimir saída

- **Extensibilidade para novos comandos**
  - Para criar um comando novo:
    1. Adicionar string no slice `commands` (em `infra/commands/commands.go`).
    2. Adicionar `case` correspondente em `Processor.Process`.
    3. Criar handler em `infra/processor/handlers/<novo_comando>.go`.

---

## Fluxos de uso e testes

- **Binário principal**: `task-cli`

- **Fluxos principais (manuais ou integração)**
  - Add:
    - `task-cli add --description "Estudar Go"` ou `task-cli add -d "Estudar Go"`
  - List:
    - `task-cli list`
  - Update:
    - `task-cli update --id 1 --description "Estudar Go (capítulo 2)"`
  - Mark done:
    - `task-cli mark-done --id 1`
  - Get:
    - `task-cli get --id 1`
  - Delete:
    - `task-cli delete --id 1`

- **Testes negativos**
  - Exemplos:
    - `task-cli add`
    - `task-cli update --id 1`
    - `task-cli get`
    - `task-cli delete`
    - `task-cli mark-done`
  - Comportamento esperado:
    - Mensagem de erro clara.
    - Ajuda com `Usage: task-cli <command> <args...>`.
    - Exit code 1.

- **Testes automatizados sugeridos**
  - Rodar:
    ```bash
    go test ./...
    ```
  - Esperado:
    - Todos os testes passam (especialmente em `core/services`).

---

## Como AIs / Skills devem usar este contexto

- **Arquivos-chave para ler primeiro**
  - `infra/commands/commands.go` → parsing de comandos, ajuda, opções.
  - `infra/processor/processor.go` → fluxo principal e roteamento para handlers.
  - `infra/processor/handlers/*.go` → regras de validação e mensagens ao usuário.
  - `core/services` (quando disponível) → regras de negócio.

- **Ao implementar mudanças**
  - **Novos comandos**:
    - Seguir o padrão: `commands` slice → `Processor.Process` → novo handler.
    - Manter validação e mensagens consistentes com os comandos existentes.
  - **Mudanças na UX de CLI**:
    - Garantir que mensagens de erro continuam claras e que `help` ainda é acionado corretamente.
  - **Testes**:
    - Atualizar ou criar testes (unitários/integrados) para novos fluxos.
    - Confirmar que `go test ./...` continua passando.

- **Linguagem**
  - A documentação e mensagens podem aparecer em PT e EN.
  - Quando em dúvida, manter os nomes de comandos e opções em inglês (ex.: `description`, `status`) e mensagens legíveis para usuários PT/EN.

---

## Pedidos comuns para a AI (sugestões)

- PT: “Adicionar um novo comando `X` seguindo o padrão dos handlers atuais, com validação e mensagens consistentes.”  
  EN: “Add a new `X` command following the existing handler pattern, including validation and user-friendly messages.”

- PT: “Refatorar o parsing de opções preservando compatibilidade com `--description` e `-d`.”  
  EN: “Refactor option parsing while keeping compatibility with both `--description` and `-d`.”

- PT: “Criar testes de integração para o fluxo completo `add → list → update → mark-done → get → delete` usando `task-cli`.”  
  EN: “Create integration tests for the full flow `add → list → update → mark-done → get → delete` using `task-cli`.”

