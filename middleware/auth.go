package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println("Validate token")
	c.Next()
}
