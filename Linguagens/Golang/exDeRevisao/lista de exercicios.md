# 🟢 1️⃣ Variáveis + If (bem simples)

Crie um programa que:

* Declare uma variável `idade`
* Se idade >= 18 → imprima `"Maior de idade"`
* Senão → `"Menor de idade"`

Só isso.

---

# 🟢 2️⃣ For (loop básico)

Faça um `for` que:

* Imprima números de 1 até 10

Depois modifica para:

* Imprimir apenas números pares

---

# 🟢 3️⃣ Função

Crie uma função:

```go
func dobro(numero int) int
```

Que retorne o dobro do número.

No `main`, chame a função e imprima o resultado.

---

# 🟢 4️⃣ Slice

Crie um slice de `int` com 5 números.

Depois:

* Percorra com `for`
* Some todos os valores
* Imprima o total

---

# 🟢 5️⃣ Struct

Crie:

```go
type Pessoa struct {
	Nome string
	Idade int
}
```

No `main`:

* Crie uma pessoa
* Imprima nome e idade

---

# 🟢 6️⃣ Método (aqui começa a ficar interessante)

Adicione um método à `Pessoa`:

```go
func (p Pessoa) Apresentar() string
```

Que retorne:

```
"Oi, eu sou <Nome> e tenho <Idade> anos"
```

Chame no `main`.

---

# 🔥 7️⃣ Exercício GERAL (misturando tudo)

Crie:

```go
type Produto struct {
	Nome  string
	Preco float64
}
```

Agora:

* Crie um slice de produtos
* Adicione 3 produtos
* Percorra com `for`
* Se preço > 100 → imprima "Produto caro"
* Senão → "Produto acessível"

