package repository

import (
	"database/sql"
	"fmt"
	"myGoApi/model"
)

type ProductRepositoy struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepositoy {
	return ProductRepositoy{
		connection: connection,
	}
}

func (pr *ProductRepositoy) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.NAME,
			&productObj.VALOR,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)

	}

	rows.Close()
	return productList, nil
}
