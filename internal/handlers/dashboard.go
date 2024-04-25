package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashboardHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	}
}

func DashboardPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/results")
	}
}

func ResultsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "results.html", gin.H{})
	}
}
