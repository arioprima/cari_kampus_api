package routes

import (
	handlers "github.com/arioprima/cari_kampus_api/handlers/wilayah"
	"github.com/arioprima/cari_kampus_api/middlewares"
	repositories "github.com/arioprima/cari_kampus_api/repositories/wilayah"
	services "github.com/arioprima/cari_kampus_api/services/wilayah"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupWilayahRoutes(route *gin.Engine, db *gorm.DB) {
	wilayahRepository := repositories.NewRepositoryWilayahImpl(db)
	wilayahService := services.NewServiceResultImpl(wilayahRepository)
	wilayahHandler := handlers.NewHandlerResultWilayah(wilayahService)

	groupRoute := route.Group("/api")
	groupRoute.GET("/wilayah", middlewares.AuthMiddleware("user"), wilayahHandler.ResultWilayahHandler)
}
