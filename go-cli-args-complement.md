### Exemplo prático com `os.Args`

Um programa que apenas imprime os argumentos recebidos:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Nome do programa:", os.Args[0])
	fmt.Println("Quantidade de args:", len(os.Args)-1)

	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg %d: %s\n", i+1, arg)
	}

	if len(os.Args) > 1 {
		first := os.Args[1]
		fmt.Println("Primeiro arg:", first)
	} else {
		fmt.Println("Nenhum argumento informado")
	}
}
```

- `os.Args` é um `[]string` com tudo que veio na linha de comando.
- `os.Args[0]` é o nome/caminho do programa.
- `os.Args[1:]` são os argumentos reais.
- Sempre cheque `len(os.Args)` antes de acessar índices como `os.Args[1]` para evitar panic.

---

### Exemplo prático com o pacote `flag`

Um programa que usa flags nomeadas (`--name`, `--age`, `--verbose`) e ainda aceita argumentos posicionais restantes:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Mundo", "nome a ser saudado")
	age := flag.Int("age", 0, "idade da pessoa")
	verbose := flag.Bool("verbose", false, "modo verboso")

	flag.Parse()

	rest := flag.Args()

	fmt.Printf("Olá, %s!\n", *name)
	if *age > 0 {
		fmt.Printf("Você tem %d anos.\n", *age)
	}
	if *verbose {
		fmt.Println("Modo verboso ligado.")
	}
	fmt.Println("Argumentos restantes:", rest)
}
```

- Use `flag.String`, `flag.Int`, `flag.Bool` para declarar flags.
- Sempre chame `flag.Parse()` antes de usar os valores.
- As funções de `flag` retornam ponteiros; use `*name`, `*age`, `*verbose`.
- `flag.Args()` retorna os argumentos que sobraram (não reconhecidos como flags).

---

### Boas práticas e exercícios sugeridos

- **Validação de entrada**:
  - Se um argumento/flag é obrigatório, verifique após `flag.Parse()` ou ao inspecionar `os.Args`.
  - Em caso de erro, mostre uma mensagem de uso clara e encerre com código diferente de zero.

- **Mensagem de ajuda simples**:

```go
if *name == "" {
	fmt.Println("Uso: myprog --name=<nome> [--age=<idade>] [--verbose] [args...]")
	flag.PrintDefaults()
	os.Exit(1)
}
```

- **Separar parsing da lógica**:
  - Mantenha o código que lê/parsa argumentos em uma função (ou no início do `main`).
  - Deixe a lógica de negócio em funções próprias para facilitar testes e reuso.

**Exercícios práticos:**

1. **Soma de números via CLI**  
   - Use `os.Args[1:]` + `strconv.Atoi` para somar inteiros.  
   - Ex.: `go run main.go 10 20 5` → imprime `35`.

2. **Conversor de temperatura**  
   - Flags: `--celsius` (float64) e `--to` (`"F"` ou `"K"`).  
   - Ex.: `go run main.go --celsius=25 --to=F`.

3. **Mini gerenciador de tarefas (em memória)**  
   - `--add "<tarefa>"` para adicionar.  
   - `--list` para listar.  
   - Comece apenas imprimindo em memória; depois você pode evoluir para salvar em arquivo.

