package main

import (
	"Univerte/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := make([]handlers.User, 10, 10)

	r := gin.Default()
	r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.IndexHandler())
	r.GET("/login", handlers.LoginHandler())
	r.GET("/register", handlers.RegisterHandler())
	r.GET("/dashboard", handlers.DashboardHandler())
	r.GET("/results", handlers.ResultsHandler())

	r.POST("/login", handlers.LoginPostHandler(&DB))
	r.POST("/register", handlers.RegisterPostHandler(&DB))
	r.POST("/dashboard", handlers.DashboardPostHandler())

	r.Run("0.0.0.0:8080")
}
