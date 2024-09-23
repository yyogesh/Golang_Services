package routes

import (
	"database/sql"
	"job_portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	{
		// AUTH ROUTES
		r.POST("/login", handlers.LoginHandler(db))
		r.POST("/register", handlers.RegisterHandler(db))
	}
}
