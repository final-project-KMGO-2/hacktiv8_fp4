package middleware

import (
	"fmt"
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Invalid Token", "Token is invalid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Failed to process request", "Error validating token", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		if !token.Valid {
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		userRole, err := jwtService.GetRoleByToken(string(authHeader))
		if err != nil || (userRole != "admin" && userRole != role) {
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		userID, err := jwtService.GetUserIDByToken(string(authHeader))
		if err != nil {
			response := common.BuildErrorResponse("Failed to process request", "Failed to get userID", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		fmt.Println(userID)
		c.Set("userID", userID)
		c.Next()
	}
}
