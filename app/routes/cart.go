package routes

import (
	"nexmedis-technical-test/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.RouterGroup, cartHandler *handler.CartHandler) {
	cartGroup := router.Group("/carts")
	{
		cartGroup.POST("/add", cartHandler.AddToCart)
		cartGroup.GET("/", cartHandler.GetCartItems)
		cartGroup.POST("/checkout", cartHandler.Checkout)
	}
}
