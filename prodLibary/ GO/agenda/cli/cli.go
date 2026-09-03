package cli

import (
	"fmt"

	"agenda/funcy"
)

func CliInit() {
	var choices int
	fmt.Println(`
 █████╗  ██████╗ ███████╗███╗   ██╗██████╗  ██████╗██╗     ██╗
██╔══██╗██╔════╝ ██╔════╝████╗  ██║██╔══██╗██╔════╝██║     ██║
███████║██║  ███╗█████╗  ██╔██╗ ██║██║  ██║██║     ██║     ██║
██╔══██║██║   ██║██╔══╝  ██║╚██╗██║██║  ██║██║     ██║     ██║
██║  ██║╚██████╔╝███████╗██║ ╚████║██████╔╝╚██████╗███████╗██║
╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═══╝╚═════╝  ╚═════╝╚══════╝╚═╝
	`)
	fmt.Println("Utilize sua lista de contatos acrescente, liste, edite e apague escolha uma opção:")

	fmt.Println("1 - Adicionar contato")
	fmt.Println("2 - Listar contatos")
	fmt.Println("3 - Editar contato")
	fmt.Println("4 - Remover contato")
	fmt.Println("5 - Sair")
	fmt.Print("->")

	fmt.Scan(&choices)

	switch choices {
	case 1:
		funcy.AdicionarContatos()
	case 2:
		funcy.ListarContatos()
	case 3:
		funcy.EditarContatos()
	case 4:
		funcy.RemoverContatos()
	case 5:
		funcy.Sair()
	default:
		fmt.Println("Opção inválida!")
	}
}
