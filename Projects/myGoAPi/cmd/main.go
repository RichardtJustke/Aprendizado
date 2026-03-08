package main

import (
	"myGoApi/controller"
	"myGoApi/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	ProductUseCase := usecase.NewProductUseCase()
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
