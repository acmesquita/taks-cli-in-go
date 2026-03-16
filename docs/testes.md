## Demonstração de testes

Abaixo alguns fluxos de teste manuais (e que também podem ser automatizados em testes de integração) usando o binário `task-cli`.

### Fluxos principais

- **Adicionar uma nova tarefa**
  - Comando:
    ```bash
    task-cli add --description "Estudar Go"
    # ou
    task-cli add -d "Estudar Go"
    ```
  - Resultado esperado:
    - Saída iniciando com `"Adding task"`.
    - Impressão do ID gerado, descrição e status (por exemplo: `1 Estudar Go PENDING`).
    - Mensagem `"Task added successfully"`.

- **Listar tarefas existentes**
  - Comando:
    ```bash
    task-cli list
    ```
  - Resultado esperado:
    - Saída `"Listing tasks"`.
    - Lista de tarefas com `ID Description Status`.
    - Mensagem `"Tasks listed successfully"`.

- **Atualizar uma tarefa existente**
  - Comando (usando o ID retornado no passo anterior):
    ```bash
    task-cli update --id 1 --description "Estudar Go (capítulo 2)"
    # ou
    task-cli update --id 1 -d "Estudar Go (capítulo 2)"
    ```
  - Resultado esperado:
    - `"Updating task"`.
    - Nova descrição impressa junto com o ID e status.
    - `"Task updated successfully"`.

- **Marcar uma tarefa como concluída**
  - Comando:
    ```bash
    task-cli mark-done --id 1
    ```
  - Resultado esperado:
    - `"Marking task as done"`.
    - Tarefa com status atualizado (ex.: `DONE`).
    - `"Task marked as done successfully"`.

- **Obter detalhes de uma tarefa específica**
  - Comando:
    ```bash
    task-cli get --id 1
    ```
  - Resultado esperado:
    - `"Getting task"`.
    - Impressão do ID, descrição e status atual.
    - `"Task found successfully"`.

- **Excluir uma tarefa**
  - Comando:
    ```bash
    task-cli delete --id 1
    ```
  - Resultado esperado:
    - `"Deleting task"`.
    - Impressão da tarefa removida.
    - `"Task deleted successfully"`.

### Testes negativos (validação de erros)

Executar comandos sem os parâmetros obrigatórios deve:

- Imprimir uma mensagem de erro clara.
- Exibir a ajuda com `Usage: task-cli <command> <args...>` e lista de comandos.
- Sair com código de erro (1).

Exemplos:

```bash
task-cli add
task-cli update --id 1
task-cli get
task-cli delete
task-cli mark-done
```

### Testes automatizados (sugestão)

Caso existam testes automatizados (por exemplo, em `core/services`), a documentação pode recomendar:

- **Execução de testes automáticos**
  ```bash
  go test ./...
  ```
  - Resultado esperado:
    - Todos os testes passando, garantindo:
      - Criação, atualização, listagem e remoção de tarefas.
      - Comportamento correto de `TaskService` ao receber entradas válidas/inválidas.

---

## Test demonstration (EN)

Below are manual test flows (that can also be turned into integration tests) using the `task-cli` binary.

### Main flows

- **Add a new task**
  - Command:
    ```bash
    task-cli add --description "Study Go"
    # or
    task-cli add -d "Study Go"
    ```
  - Expected result:
    - Output starting with `"Adding task"`.
    - Printed ID, description and status (e.g. `1 Study Go PENDING`).
    - `"Task added successfully"`.

- **List existing tasks**
  - Command:
    ```bash
    task-cli list
    ```
  - Expected result:
    - `"Listing tasks"`.
    - List of tasks with `ID Description Status`.
    - `"Tasks listed successfully"`.

- **Update an existing task**
  - Command (using the ID from the previous step):
    ```bash
    task-cli update --id 1 --description "Study Go (chapter 2)"
    # or
    task-cli update --id 1 -d "Study Go (chapter 2)"
    ```
  - Expected result:
    - `"Updating task"`.
    - Updated description printed with ID and status.
    - `"Task updated successfully"`.

- **Mark a task as done**
  - Command:
    ```bash
    task-cli mark-done --id 1
    ```
  - Expected result:
    - `"Marking task as done"`.
    - Task printed with updated status (e.g. `DONE`).
    - `"Task marked as done successfully"`.

- **Get details of a specific task**
  - Command:
    ```bash
    task-cli get --id 1
    ```
  - Expected result:
    - `"Getting task"`.
    - Task ID, description and current status.
    - `"Task found successfully"`.

- **Delete a task**
  - Command:
    ```bash
    task-cli delete --id 1
    ```
  - Expected result:
    - `"Deleting task"`.
    - Deleted task printed.
    - `"Task deleted successfully"`.

### Negative tests (error validation)

Running commands without required parameters should:

- Print a clear error message.
- Show the help text with `Usage: task-cli <command> <args...>` and the commands list.
- Exit with error code 1.

Examples:

```bash
task-cli add
task-cli update --id 1
task-cli get
task-cli delete
task-cli mark-done
```

### Automated tests (suggestion)

If you have automated tests (for example, in `core/services`), the documentation can recommend:

- **Run automated tests**
  ```bash
  go test ./...
  ```
  - Expected result:
    - All tests pass, validating:
      - Task creation, update, listing and deletion.
      - Correct `TaskService` behavior for valid and invalid inputs.

