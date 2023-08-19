package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-like-twitter-api/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Signup(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err != nil {
		// c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		// c.Abort()
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if err := createUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "hello from the Register Point",
	})
}
