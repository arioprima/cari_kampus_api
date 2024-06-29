package middlewares

import (
	repositories "github.com/arioprima/cari_kampus_api/repositories/role"
	"github.com/arioprima/cari_kampus_api/schemas"
	"net/http"
	"strings"

	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/pkg"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		var roleUser string = userRole
		// Error response structure
		errorResponse := schemas.SchemaErrorResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
			Error:  "Authorization is required for this endpoint",
		}

		// Check Authorization header
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
			return
		}
		token = fields[1]

		// Load configuration and open database connection
		loadConfig, err := config.LoadConfig(".")
		if err != nil {
			errorResponse = schemas.SchemaErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Error:  "Failed to load configuration",
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
			return
		}

		db, err := config.OpenConnection(&loadConfig)
		if err != nil {
			errorResponse = schemas.SchemaErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Error:  "Failed to open database connection",
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
			return
		}

		// Validate token and get user role
		sub, err := pkg.ValidateToken(token, loadConfig.TokenSecret)
		if err != nil {
			errorResponse = schemas.SchemaErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Error:  "Invalid token",
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
			return
		}

		userRole, ok := sub.(map[string]interface{})["role_id"].(string)
		if !ok {
			errorResponse = schemas.SchemaErrorResponse{
				Code:   http.StatusForbidden,
				Status: "Forbidden",
				Error:  "Role ID not found in token claims",
			}
			ctx.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
			return
		}

		// Get role name by role ID
		roleName, err := repositories.GetRoleNameById(userRole, db)
		if err != nil {
			errorResponse = schemas.SchemaErrorResponse{
				Code:   http.StatusNotFound,
				Status: "Not Found",
				Error:  "Role not found",
			}
			ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse)
			return
		}

		if roleUser == "admin" {
			if roleName == roleUser {
				ctx.Next()
			} else {
				errorResponse = schemas.SchemaErrorResponse{
					Code:   http.StatusForbidden,
					Status: "Forbidden",
					Error:  "You are not authorized to access this endpoint",
				}
				ctx.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
			}
		} else {
			ctx.Next()
		}
	}
}
