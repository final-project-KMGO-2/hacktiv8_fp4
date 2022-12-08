package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddNewProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	UpdateProductByID(ctx *gin.Context)
	DeleteProductById(ctx *gin.Context)
}

// TODO: buy new product 

type productController struct {
	jwtService service.JWTService
	productService service.ProductService
}

func NewProductController(jt service.JWTService, ps service.ProductService) ProductController {
	return &productController{jwtService: jt, productService: ps}
}

func (ps *productController) AddNewProduct(ctx *gin.Context) {
	productCreate := entity.ProductCreate{}
	err := ctx.ShouldBind(&productCreate)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := ps.productService.CreateNewProduct(ctx, productCreate)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "Product Created", data)
	ctx.JSON(http.StatusCreated, response)
}
func (ps *productController) GetAllProducts(ctx *gin.Context) {
	data, err := ps.productService.GetAllProducts(ctx.Request.Context())
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusCreated, response)
}

func (ps *productController) UpdateProductByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productUpdate := entity.ProductUpdate{}
	err = ctx.ShouldBind(&productUpdate)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := ps.productService.UpdateProductById(ctx.Request.Context(), productUpdate, uint64(id))

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (ps *productController) DeleteProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = ps.productService.DeleteProductById(ctx.Request.Context(), uint64(id))
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Product Deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
