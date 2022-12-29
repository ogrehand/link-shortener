package middleware

import (
	"net/http"
	"shortenerBE/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, ok := c.Request.Header["Authorization"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "please login"})
			c.Abort()
			return
		}
		token := strings.Split(auth[0], " ")

		_, err := model.CheckTokenRedis(token[1])
		if err == redis.Nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "please login"})
			c.Abort()
			return
		}

		// before request

		// c.Next()

	}
}
