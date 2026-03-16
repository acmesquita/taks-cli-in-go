## Soluções

Alguns pontos de solução adotados pela implementação.

- **Validação antecipada de comandos e opções**
  - `ParseCommands` e `validateCommands` garantem que apenas comandos suportados sejam executados.
  - `ParseOptions` converte os argumentos em um mapa genérico, suportando tanto `--description` quanto `-d`.
  - Para `add` e `update`, há validações específicas em `ParseOptions` para garantir que `description` não seja vazia, evitando que o fluxo siga com dados inválidos.

- **Tratamento consistente de erros de entrada**
  - Em vários handlers (`AddTask`, `UpdateTask`, `FindTask`, `DeleteTask`, `MarkDoneTask`):
    - Quando um parâmetro obrigatório falta (`id` ou `description`), o programa:
      - Imprime uma mensagem clara (ex.: `"Description is required"`, `"ID is required"`).
      - Chama `commands.HandleHelperMessage` para exibir a ajuda.
      - Finaliza com `os.Exit(1)` indicando erro de uso.

- **Encapsulamento da lógica de negócio**
  - Toda a lógica de tarefas (criar, atualizar, listar, excluir, marcar como concluída) está dentro de `TaskService`, mantendo a CLI e os handlers finos, focados apenas em:
    - Validar entrada
    - Orquestrar chamadas para o serviço
    - Exibir resultados

- **Feedback explícito ao usuário em cada operação**
  - Cada handler imprime uma mensagem de contexto (por exemplo, `"Adding task"`, `"Updating task"`, `"Listing tasks"`).
  - Após a operação, imprime os detalhes da tarefa e uma mensagem de sucesso (`"Task added successfully"`, `"Task deleted successfully"`, etc.).
  - Isso facilita o entendimento do fluxo, especialmente em cenários de linha de comando.

- **Extensibilidade para novos comandos**
  - Novos comandos podem ser adicionados:
    - Incluindo o comando no slice `commands`.
    - Adicionando um novo *case* no `Processor.Process`.
    - Criando um novo handler especializado.
  - Esse padrão favorece a evolução incremental da ferramenta.

---

## Solutions (EN)

Key solution choices and behaviors.

- **Early validation of commands and options**
  - `ParseCommands` and `validateCommands` ensure only supported commands are executed.
  - `ParseOptions` converts CLI arguments into a generic map, supporting both `--description` and `-d`.
  - For `add` and `update`, there are specialized validations in `ParseOptions` that guarantee `description` is not empty, preventing invalid flows.

- **Consistent handling of input errors**
  - In multiple handlers (`AddTask`, `UpdateTask`, `FindTask`, `DeleteTask`, `MarkDoneTask`):
    - When a required parameter is missing (`id` or `description`), the program:
      - Prints a clear error message (`"Description is required"`, `"ID is required"`).
      - Calls `commands.HandleHelperMessage` to display usage help.
      - Exits with `os.Exit(1)` to signal incorrect usage.

- **Business logic encapsulation**
  - All task logic (create, update, list, delete, mark as done) is within `TaskService`, keeping the CLI and handlers thin and focused on:
    - Input validation
    - Orchestration of service calls
    - Rendering output

- **Explicit user feedback for every operation**
  - Each handler prints a contextual message (e.g., `"Adding task"`, `"Updating task"`, `"Listing tasks"`).
  - After the operation, it prints task details and a success message (`"Task added successfully"`, `"Task deleted successfully"`, etc.).
  - This improves usability when running in a terminal.

- **Extensibility for new commands**
  - New commands can be added by:
    - Including the command string in the `commands` slice.
    - Adding a new `case` in `Processor.Process`.
    - Implementing a new specialized handler.
  - This pattern supports incremental evolution of the tool.

