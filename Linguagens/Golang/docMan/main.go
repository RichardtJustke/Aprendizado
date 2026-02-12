package main

import "fmt"

type Usuario struct {
	ID    int
	Nome  string
	Ativo bool
}

func main() {
	u := Usuario{ID: 1, Nome: "Richardt", Ativo: true}
	fmt.Println(u.ID)
}
