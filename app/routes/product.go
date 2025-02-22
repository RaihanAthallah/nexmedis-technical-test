package routes

import (
	"nexmedis-technical-test/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, productHandler *handler.ProductHandler) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/:id", productHandler.GetProductByID)
		productGroup.GET("", productHandler.GetProducts)
	}
}
