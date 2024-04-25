package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}
}

func RegisterPostHandler(DB *[]User) gin.HandlerFunc {

	return func(c *gin.Context) {
		var a User = User{}
		a.Email = c.PostForm("email")
		a.Password = c.PostForm("password")
		a.Firstname = c.PostForm("firstname")
		a.Surname = c.PostForm("surname")

		fmt.Println(a)

		_ = append(*DB, a)

		c.Redirect(http.StatusSeeOther, "/login")
	}
}
