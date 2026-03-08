package main

import (
	"myGoApi/controller"
	"myGoApi/db"
	"myGoApi/repository"
	"myGoApi/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//Camada do repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada do usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controladores
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")

}
