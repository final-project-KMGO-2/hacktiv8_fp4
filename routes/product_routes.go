package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"

	"github.com/gin-gonic/gin"
)

func GenerateProductRoutes(router *gin.Engine, productController controller.ProductController, jwtService middleware.JWTService) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", middleware.Authenticate(jwtService, "member"), productController.GetAllProducts)
		productRoutes.POST("", middleware.Authenticate(jwtService, "admin"), productController.AddNewProduct)
		productRoutes.PUT("/:productId", middleware.Authenticate(jwtService, "admin"), productController.UpdateProductByID)
		productRoutes.DELETE("/:productId", middleware.Authenticate(jwtService, "admin"), productController.DeleteProductById)
	}
}
