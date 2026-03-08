package controller

import (
	"myGoApi/model"
	"myGoApi/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)
type productController struct{
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController{
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context){

	products := []model.Product{
		{
			ID: 1,
			NAME: "batata frita",
			VALOR: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
