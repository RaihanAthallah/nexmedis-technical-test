package routes

import (
	"nexmedis-technical-test/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupBankRoutes(router *gin.RouterGroup, bankHandler *handler.BankHandler) {
	bankGroup := router.Group("/banks")
	{
		bankGroup.POST("/deposit", bankHandler.Deposit)
		bankGroup.POST("/withdraw", bankHandler.Withdraw)
	}
}
