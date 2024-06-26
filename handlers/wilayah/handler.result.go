package handlers

import (
	"github.com/arioprima/cari_kampus_api/helpers"
	services "github.com/arioprima/cari_kampus_api/services/wilayah"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerResult struct {
	service services.ServiceResult
}

func NewHandlerResultWilayah(service services.ServiceResult) *HandlerResult {
	return &HandlerResult{service: service}
}

func (h *HandlerResult) ResultWilayahHandler(ctx *gin.Context) {
	res, err := h.service.ResultWilayahService()

	switch err.Type {
	case "error_01":
		helpers.ApiResponse(ctx, http.StatusNotFound, "error", "Data not found", nil, nil)
	default:
		helpers.ApiResponse(ctx, http.StatusOK, "success", "Result Get Data Wilayah Successfully", res, nil)
	}
}
