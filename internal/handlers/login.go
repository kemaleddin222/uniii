package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Firstname string
	Surname   string
	Password  string
	Email     string
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func LoginPostHandler(DB *[]User) gin.HandlerFunc {

	return func(c *gin.Context) {

		var found bool = false
		for _, u := range *DB {
			if u.Password == c.PostForm("password") && u.Email == c.PostForm("email") {
				fmt.Println("User found in in-memory database !!!")
				found = true
				break
			}
		}

		if !found {
			c.HTML(http.StatusNotFound, "Error occured !", gin.H{})
		}
		c.Redirect(http.StatusSeeOther, "/dashboard")
	}
}
