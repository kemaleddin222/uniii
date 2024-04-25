package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{})
	}
}
