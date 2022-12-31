package middleware

import (
	"fmt"
	"net/http"
	"shortenerBE/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Set("abort", true)
		// c.Abort()
		auth, ok := c.Request.Header["Authorization"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "please login 1"})
			return
		}

		token := strings.Split(auth[0], " ")

		val, err := model.CheckTokenRedis(token[1])
		if err == redis.Nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "please login 2"})
			return
		}
		// fullPath := c.FullPath()
		method := c.Request.Method
		// match1, err := regexp.MatchString("/logout", fullPath)
		path := strings.Split(c.FullPath(), "/")
		fmt.Println(path, method)
		if path[1] == "users" {
			if path[len(path)-1] == "logout" {
				return
			}
			if val != path[len(path)-1] {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Message": "you have no right to do this"})
				return
			}
		}

		c.Next()

	}
}
