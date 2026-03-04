package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade string
}

func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Oi me chamo %s e tenho %i anos", p.Nome, p.Idade)
}
func main() {
	p := Pessoa{
		Nome:  "Richardt",
		Idade: "21",
	}
	fmt.Println(p.Saudacao())
}
