package auth

import (
	"bytes"
	"io"
	"job_portal/internal/models"
	"job_portal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func PasswordValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the request body
		bodyBytes, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(400, gin.H{"error": "Error reading request body"})
			c.Abort()
			return
		}
		// Create new reader with the bytes
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Parse the request
		var req models.ChangePasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}
		// Restore the request body for the next middleware/handler
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		isValid, errors := utils.ValidatePasswordStrength(req.NewPassword)
		if !isValid {
			c.JSON(400, gin.H{
				"error":   "Password validation failed",
				"details": errors,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
