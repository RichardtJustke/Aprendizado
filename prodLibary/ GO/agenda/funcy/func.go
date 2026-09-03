package funcy

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	NomeContact     string `json:"nome"`
	TelefoneContact int    `json:"telefone"`
	EmailContact    string `json:"email"`
}

func AdicionarContatos() {
	var contato User
	var contatos []User

	fmt.Println("Digite o nome do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.NomeContact)
	fmt.Println("Digite o telefone do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.TelefoneContact)
	fmt.Println("Digite o email do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.EmailContact)

	contatos = append(contatos, contato)

	dadosJSON, err := json.MarshalIndent(contatos, "", " ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("data/contacts.json", dadosJSON, 0o644)
	if err != nil {
		panic(err)
	}
}

func ListarContatos() {
}

func EditarContatos() {
}

func RemoverContatos() {
}

func Sair() {
}
