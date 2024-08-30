package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RequestIDHeader = "X-Request-ID"
	ServiceIDHeader = "X-Service-ID"
)

func RequestID() gin.HandlerFunc {
	return checkHeader(RequestIDHeader)
}

func ServiceID() gin.HandlerFunc {
	return checkHeader(ServiceIDHeader)
}

func checkHeader(headerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceID := c.GetHeader(headerName)
		if serviceID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("'%s' header is required", headerName)})
			c.Abort()
			return
		}
		c.Next()
	}
}

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth {
			if user == "foo" && password == "bar" {
				c.Next()
				return
			}
		}

		c.AbortWithStatus(401)
	}
}
