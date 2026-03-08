package controller

import (
	"myGoApi/model"
	"net/http"
	"github.com/gin-gonic/gin"
)
type productController struct{
	//usecase
}

func NewProductController() productController{
	return productController{}
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
