package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHistoryController interface {
	CreateTransactionHistory(ctx *gin.Context)
	GetAllTransactionHistory(ctx *gin.Context)
	GetTransactionHistoryByUserID(ctx *gin.Context)
}

type transactionHistoryController struct {
	jwtService                middleware.JWTService
	transactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryController(ths service.TransactionHistoryService, js middleware.JWTService) TransactionHistoryController {
	return &transactionHistoryController{
		transactionHistoryService: ths,
		jwtService:                js,
	}
}

func (c *transactionHistoryController) CreateTransactionHistory(ctx *gin.Context) {
	var transactionHistory entity.TransactionHistoryCreate

	errBind := ctx.ShouldBind(&transactionHistory)
	if errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userID, ok := ctx.MustGet("userID").(uint64)
	if !ok {
		response := common.BuildErrorResponse("Failed to get transaction history", "userID not found", common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	transactionHistory.UserID = userID
	result, err := c.transactionHistoryService.CreateTransactionHistory(ctx.Request.Context(), transactionHistory)
	if err != nil {
		response := common.BuildErrorResponse("Failed to create transaction history", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *transactionHistoryController) GetAllTransactionHistory(ctx *gin.Context) {
	result, err := c.transactionHistoryService.GetAllTransactionHistory(ctx.Request.Context())
	if err != nil {
		response := common.BuildErrorResponse("Failed to get transaction history", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *transactionHistoryController) GetTransactionHistoryByUserID(ctx *gin.Context) {
	userID, ok := ctx.MustGet("userID").(uint64)
	if !ok {
		response := common.BuildErrorResponse("Failed to get transaction history", "userID not found", common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result, err := c.transactionHistoryService.GetTransactionHistoryByUserID(ctx.Request.Context(), userID)
	if err != nil {
		response := common.BuildErrorResponse("Failed to get transaction history", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusCreated, response)
}
