package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/topup", middleware.Authenticate(jwtService, "member"), userController.UpdateUserBalance)
		userRoutes.DELETE("/delete-account", middleware.Authenticate(jwtService, "member"), userController.DeleteUser)
	}
}
