package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "admin" && password == "admin" {
		log.Printf("User %s authorized", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
