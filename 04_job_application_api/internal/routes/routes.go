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
		r.GET("/jobs", handlers.GetAllJobsHandler(db))
		r.POST("forgotpassword", handlers.ForgotPasswordHandler(db))

		// User routes // employer
		authenticated := r.Group("/")
		authenticated.Use(auth.AuthMiddleware())
		authenticated.GET("/users/:id", handlers.GetUserByIdHandler(db))
		authenticated.PUT("/users/:id", handlers.UpdateUserProfileHandler(db))
		authenticated.POST("/users/:id/picture", handlers.UpdateUserProfilePcitureHandler(db))
		authenticated.PUT("users/change-password", auth.PasswordValidationMiddleware(), handlers.ChangePasswordHandler(db))

		// job routes
		authenticated.POST("/jobs", handlers.CreateJobHandler(db))
		authenticated.GET("/jobsByUser", handlers.GetAllJobsByUserHandler(db))
		authenticated.GET("/jobs/:id", handlers.GetJobByIdHandler(db))
		authenticated.PUT("/jobs/:id", handlers.UpdateJobByHandler(db))
		authenticated.DELETE("/jobs/:id", handlers.DeleteJobByHandler(db))

		authenticated.GET("/users", handlers.GetAllUsersHandler(db))
		authenticated.DELETE("/users/:id", handlers.DeleteUserByIdHandler(db))
	}
}
