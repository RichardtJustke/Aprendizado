package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	p := Pessoa{Nome: "Richardt", Idade: 21}
	fmt.Println(p.Nome, p.Idade)
}
