package controller

import (
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddNewProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	UpdateProductByID(ctx *gin.Context)
	DeleteProductById(ctx *gin.Context)
}


type productController struct {
	jwtService service.JWTService
	productService service.ProductService
}

func NewProductController(jt service.JWTService, ps service.ProductService) ProductController {
	return &productController{jwtService: jt, productService: ps}
}

func (ps *productController) AddNewProduct(ctx *gin.Context) {}
func (ps *productController) GetAllProducts(ctx *gin.Context) {}
func (ps *productController) UpdateProductByID(ctx *gin.Context) {}
func (ps *productController) DeleteProductById(ctx *gin.Context) {}
