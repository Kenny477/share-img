package main

import (
	"net/http"

	"github.com/Kenny477/share-img/controllers"
	"github.com/Kenny477/share-img/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Test upload page
	r.LoadHTMLGlob("views/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.tmpl", gin.H{})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// File upload and retrieve
	file := r.Group("/file")
	{
		file.POST("/", controllers.Upload)
		file.GET("/:id", controllers.Retrieve)
	}

	return r
}

func main() {
	db.ConnectToDB()
	db.MigrateDB()

	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
