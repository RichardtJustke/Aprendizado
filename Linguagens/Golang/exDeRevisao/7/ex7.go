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
	)
	fmt.Println(prod)
}
