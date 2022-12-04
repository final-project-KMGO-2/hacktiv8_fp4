package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func TransactionHistoryRoutes(router *gin.Engine, transactionHistoryController controller.TransactionHistoryController, jwtService service.JWTService) {
	transactionHistoryRoutes := router.Group("/transactions")
	{
		transactionHistoryRoutes.GET("/my-transactions", transactionHistoryController.GetTransactionHistoryByUserID)
		transactionHistoryRoutes.GET("/user-transactions", middleware.Authenticate(jwtService, "admin"), transactionHistoryController.GetAllTransactionHistory)
		transactionHistoryRoutes.POST("", transactionHistoryController.CreateTransactionHistory)
	}
}
