package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"

	"github.com/gin-gonic/gin"
)

func TransactionHistoryRoutes(router *gin.Engine, transactionHistoryController controller.TransactionHistoryController, jwtService middleware.JWTService) {
	transactionHistoryRoutes := router.Group("/transactions")
	{
		transactionHistoryRoutes.GET("/my-transactions", middleware.Authenticate(jwtService, "member"), transactionHistoryController.GetTransactionHistoryByUserID)
		transactionHistoryRoutes.GET("/user-transactions", middleware.Authenticate(jwtService, "admin"), transactionHistoryController.GetAllTransactionHistory)
		transactionHistoryRoutes.POST("", middleware.Authenticate(jwtService, "member"), transactionHistoryController.CreateTransactionHistory)
	}
}
