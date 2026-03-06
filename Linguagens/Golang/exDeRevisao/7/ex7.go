package main

import "fmt"

type Produto struct {
	NomeProd string
	PrecProd float64
}

func main() {

	prod := []Produto{
		{
			NomeProd: "ACER NitroV15",
			PrecProd: 8999.00,
		},
	}

	prod = append(prod, Produto{
		NomeProd: "Mackbook m4",
		PrecProd: 12999.00,
	},
		Produto{
			NomeProd: "Dell XPS",
			PrecProd: 15599.00,
		},
		Produto{
			NomeProd: "Asus RogStrix g16",
			PrecProd: 29999.00,
		},
		Produto{
			NomeProd: "Ideapad Slim 3",
			PrecProd: 1599.00,
		},
	)
	for _, produto := range prod {
		if produto.PrecProd > 2000 {
			fmt.Println(produto.NomeProd, "é muito caro")
		} else {
			fmt.Println(produto.NomeProd, "é acessivel")
		}
	}
	fmt.Println(prod[0])
}
