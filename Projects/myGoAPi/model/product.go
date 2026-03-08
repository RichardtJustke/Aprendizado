package model

type Product struct{
	ID int `json:"id_product"`
	NAME string `json:"name_product"`
	VALOR float64 `json:"price_product"`
}
