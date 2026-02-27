# 🟢 1. Estrutura básica de um programa Go

Todo programa começa assim:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

### O que é cada coisa?

* `package main` → diz que esse é um programa executável.
* `import "fmt"` → importa biblioteca de print.
* `func main()` → ponto de entrada do programa.
* `fmt.Println()` → imprime no terminal.

---

# 🟢 2. Variáveis

## Forma normal

```go
var nome string = "Richardt"
```

## Forma curta (mais usada)

```go
nome := "Richardt"
idade := 20
```

Go descobre o tipo automaticamente.

---

# 🟢 3. Tipos básicos

| Tipo    | Exemplo      |
| ------- | ------------ |
| string  | "Olá"        |
| int     | 10           |
| float64 | 3.14         |
| bool    | true / false |

---

# 🟢 4. Condições (if)

```go
idade := 18

if idade >= 18 {
	fmt.Println("Maior de idade")
} else {
	fmt.Println("Menor de idade")
}
```

Simples. Sem parênteses obrigatórios igual JS.

---

# 🟢 5. Laços (for)

Em Go só existe `for`.

## Contador

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

## Tipo while

```go
contador := 0

for contador < 5 {
	fmt.Println(contador)
	contador++
}
```

---

# 🟢 6. Funções

```go
func soma(a int, b int) int {
	return a + b
}
```

Uso:

```go
resultado := soma(5, 3)
fmt.Println(resultado)
```

---

# 🟢 7. Arrays e Slices

## Array (tamanho fixo)

```go
var numeros [3]int = [3]int{1, 2, 3}
```

## Slice (mais usado)

```go
numeros := []int{1, 2, 3}
numeros = append(numeros, 4)
```

Slice é tipo array dinâmico.

---

# 🟢 8. Struct (modelo de dados)

Aqui começa o backend de verdade.

```go
type Pessoa struct {
	Nome  string
	Idade int
}
```

Criando:

```go
p := Pessoa{
	Nome:  "Richardt",
	Idade: 20,
}
```

Acessando:

```go
fmt.Println(p.Nome)
```

Struct ≠ array
Struct = modelo com vários campos.

---

# 🟢 9. Métodos

Função ligada à struct.

```go
func (p Pessoa) Saudacao() string {
	return "Olá, meu nome é " + p.Nome
}
```

Uso:

```go
fmt.Println(p.Saudacao())
```

---

# 🟢 10. Ponteiros (nível importante)

Se quiser modificar a struct:

```go
func (p *Pessoa) FazerAniversario() {
	p.Idade++
}
```

O `*` significa que você está mexendo no original, não numa cópia.

---

# 🟢 11. Visibilidade (público e privado)

Em Go:

* Letra maiúscula → público
* Letra minúscula → privado

```go
type User struct {
	Name string
	age  int
}
```

`Name` pode ser acessado fora do pacote.
`age` não.

---

# 🟢 12. Organização de projeto

Estrutura comum:

```
meuprojeto/
│
├── go.mod
├── main.go
└── user.go
```

Criar projeto:

```bash
go mod init nome-do-projeto
```

Rodar:

```bash
go run .
```

Build:

```bash
go build
```

---

# 🧠 RESUMÃO FINAL


