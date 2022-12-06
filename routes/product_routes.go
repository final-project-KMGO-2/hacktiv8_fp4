package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func GenerateProductRoutes(router *gin.Engine, productController controller.ProductController, jwtService service.JWTService) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", middleware.Authenticate(jwtService, "member"), productController.AddNewProduct)
		productRoutes.GET("", middleware.Authenticate(jwtService, "member"), productController.GetAllProducts)
		productRoutes.PUT("/:productId", middleware.Authenticate(jwtService, "member"), productController.UpdateProductByID)
		productRoutes.DELETE("/:productId", middleware.Authenticate(jwtService, "member"), productController.DeleteProductById)
	}
}