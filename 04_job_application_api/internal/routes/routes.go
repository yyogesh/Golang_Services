package routes

import (
	"database/sql"
	"job_portal/internal/auth"
	"job_portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	{
		// AUTH ROUTES
		r.POST("/login", handlers.LoginHandler(db))
		r.POST("/register", handlers.RegisterHandler(db))

		// User routes // employer
		authenticated := r.Group("/")
		authenticated.Use(auth.AuthMiddleware())
		authenticated.GET("/users/:id", handlers.GetUserByIdHandler(db))
		authenticated.PUT("/users/:id", handlers.UpdateUserProfileHandler(db))
		authenticated.POST("/users/:id/picture", handlers.UpdateUserProfilePcitureHandler(db))
	}
}
