package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if ok && user == "admin" && password == "admin" {
		log.Printf("User %s authorized", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
