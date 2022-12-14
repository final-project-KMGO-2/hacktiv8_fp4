package controller

import (
	"fmt"
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateUserBalance(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  middleware.JWTService
}

func NewUserController(us service.UserService, as service.AuthService, js middleware.JWTService) UserController {
	return &userController{
		userService: us,
		authService: as,
		jwtService:  js,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user entity.UserRegister

	if user.Role == "" {
		user.Role = "member"
	}

	errBind := ctx.ShouldBind(&user)
	if errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isDuplicateEmail, err := c.authService.CheckEmailDuplicate(ctx.Request.Context(), user.Email)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if isDuplicateEmail {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdUser, err := c.userService.CreateUser(ctx.Request.Context(), user)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	response := common.BuildResponse(true, "OK", createdUser)
	ctx.JSON(http.StatusCreated, response)
}

func (c *userController) Login(ctx *gin.Context) {
	var userLogin entity.UserLogin
	if errBind := ctx.ShouldBind(&userLogin); errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, err := c.authService.VerifyCredential(ctx.Request.Context(), userLogin.Email, userLogin.Password)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !authResult {
		response := common.BuildErrorResponse("Error Logging in", "Invalid Credentials", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := c.userService.GetUserByEmail(ctx.Request.Context(), userLogin.Email)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := strconv.FormatUint(uint64(user.ID), 10)
	generatedToken := c.jwtService.GenerateToken(userId, user.Role)
	response := common.BuildResponse(true, "OK", generatedToken)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) UpdateUserBalance(ctx *gin.Context) {
	var userUpdateBalance entity.UserUpdateBalance
	errBind := ctx.ShouldBind(&userUpdateBalance)

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
	balance, err := c.userService.IncreaseUserBalance(ctx.Request.Context(), userID, userUpdateBalance.Balance)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update user balance", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := fmt.Sprintf("Your balance has been successfully updated to Rp%v", balance)
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	userID, ok := ctx.MustGet("userID").(uint64)
	if !ok {
		response := common.BuildErrorResponse("Failed to get transaction history", "userID not found", common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err := c.userService.DeleteUser(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete user", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Your account has been successfully deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
