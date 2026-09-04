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

	fmt.Println("Digite o nome do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.NomeContact)
	fmt.Println("Digite o telefone do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.TelefoneContact)
	fmt.Println("Digite o email do contato:")
	fmt.Print("->")
	fmt.Scan(&contato.EmailContact)

	dados, err := os.ReadFile("data/contacts.json")
	if err != nil {
		dados = []byte("[]")
	}

	var contatos []User

	err = json.Unmarshal(dados, &contatos)
	if err != nil {
		panic(err)
	}

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
	dados, err := os.ReadFile("data/contacts.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo ou nenhum contato cadastrado:", err)
		return
	}

	var contatos []User

	err = json.Unmarshal(dados, &contatos)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	if len(contatos) == 0 {
		fmt.Println("Nenhum contato encontrado.")
		return
	}

	fmt.Println("\n=== LISTA DE CONTATOS ===")
	for i, contato := range contatos {
		fmt.Printf("Contato #%d\n", i+1)
		fmt.Printf("  Nome: %s\n", contato.NomeContact)
		fmt.Printf("  Telefone: %d\n", contato.TelefoneContact)
		fmt.Printf("  Email: %s\n", contato.EmailContact)
		fmt.Println("-----------------------")
	}
}

func EditarContatos() {
	dados, err := os.ReadFile("data/contacts.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	var contatos []User
	err = json.Unmarshal(dados, &contatos)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	if len(contatos) == 0 {
		fmt.Println("Nenhum contato cadastrado para editar.")
		return
	}

	var nomeBusca string
	fmt.Print("Digite o nome do contato que deseja editar: ")
	fmt.Scan(&nomeBusca)

	indiceEncontrado := -1
	for i, contato := range contatos {
		if contato.NomeContato == nomeBusca {
			indiceEncontrado = i
			break
		}
	}

	if indiceEncontrado == -1 {
		fmt.Println("Contato não encontrado!")
		return
	}
	fmt.Println("Contato encontrado! Digite os novos dados:")

	fmt.Print("Novo Nome: ")
	fmt.Scan(&contatos[indiceEncontrado].NomeContact)

	fmt.Print("Novo Telefone: ")
	fmt.Scan(&contatos[indiceEncontrado].TelefoneContact)

	fmt.Print("Novo Email: ")
	fmt.Scan(&contatos[indiceEncontrado].EmailContact)

	novosDadosJSON, err := json.MarshalIndent(contatos, "", "  ")
	if err != nil {
		fmt.Println("Erro ao gerar JSON:", err)
		return
	}
	err = os.WriteFile("data/contacts.json", novosDadosJSON, 0o644)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
		return
	}

	fmt.Println("Contato editado com sucesso!")
}

func RemoverContatos() {
}

func Sair() {
	fmt.Println("Saindo do sistema... Até logo!")
	os.Exit(0)
}
