package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handlers "github.com/arioprima/cari_kampus_api/handlers/auth"
	repositories "github.com/arioprima/cari_kampus_api/repositories/auth"
	services "github.com/arioprima/cari_kampus_api/services/auth"
)

func SetupAuthRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	// Initialize dependencies
	loginRepository := repositories.NewRepositoryLoginImpl(db)
	loginService := services.NewServiceLoginImpl(loginRepository)
	loginHandler := handlers.NewHandlerLogin(loginService)

	// Setup routes
	groupRoute := route.Group("/api")
	groupRoute.POST("/login", loginHandler.LoginHandler)

	return route
}
