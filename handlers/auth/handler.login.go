package handlers

import (
	"github.com/arioprima/cari_kampus_api/helpers"
	"github.com/arioprima/cari_kampus_api/pkg"
	"github.com/arioprima/cari_kampus_api/schemas"
	services "github.com/arioprima/cari_kampus_api/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerLogin struct {
	Service services.ServiceLogin
}

func NewHandlerLogin(service services.ServiceLogin) *HandlerLogin {
	return &HandlerLogin{Service: service}
}

func (h *HandlerLogin) LoginHandler(ctx *gin.Context) {
	var loginRequest schemas.SchemaAuth
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", err.Error())
		return
	}

	config := []schemas.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "email",
			Message: "Email is required",
		},
		{
			Tag:     "email",
			Field:   "email",
			Message: "Email is not valid",
		},
		{
			Tag:     "required",
			Field:   "password",
			Message: "Password is required",
		},
		{
			Tag:     "min",
			Field:   "password",
			Message: "Password must be at least 3 characters",
			Value:   "3",
		},
	}
	errResponse, errCount := pkg.ValidatorLogin(&loginRequest, config)
	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", errResponse)
		return
	}

	res, err := h.Service.LoginService(ctx, &loginRequest)

	switch err.Type {
	case "error_01":
		helpers.ApiResponse(ctx, http.StatusNotFound, "error", "Email not found", nil)
		return
	case "error_02":
		helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Internal server error", nil)
		return
	case "error_03":
		helpers.ApiResponse(ctx, http.StatusUnauthorized, "error", "Password is incorrect", nil)
		return
	case "error_04":
		helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Internal server error", nil)
		return
	default:
		helpers.ApiResponse(ctx, http.StatusOK, "success", "Login successfully", res)
	}

}
