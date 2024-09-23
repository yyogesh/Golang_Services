package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Handler that returns a string
	r.GET("/hello", helloHandler)
	// Handler that returns JSON
	r.GET("/user", getUserHandler)
	// Handler that returns an array of maps
	r.GET("/items", getItemHandler)
	// Handler that returns a file
	r.GET("/download", downloadFileHanlder)
	// Handler that returns a file
	r.GET("/download/:filename", downloadableFileHanlder)

	// query parameter
	r.GET("/search", queryParamHandler)

	r.Run(":8080")
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func getUserHandler(c *gin.Context) {
	user := gin.H{
		"name": "john",
		"age":  18,
	}
	c.JSON(http.StatusOK, user)
}

func getItemHandler(c *gin.Context) {
	items := []gin.H{
		{"id": 1, "name": "Item 1"},
		{"id": 2, "name": "Item 2"},
		{"id": 3, "name": "Item 3"},
	}
	c.JSON(http.StatusOK, items)
}

func downloadFileHanlder(c *gin.Context) {
	c.File("./test.txt")
}

func downloadableFileHanlder(c *gin.Context) {
	filename := c.Param("filename")
	baseDir := "./files"
	filePath := filepath.Join(baseDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
}

func queryParamHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "Hello, %s!", name)
}
