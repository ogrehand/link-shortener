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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token Missing"})
			return
		}

		token := strings.Split(auth[0], " ")

		val, err := model.CheckTokenRedis(token[1])
		if err == redis.Nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token Expired"})
			return
		}
		// fullPath := c.FullPath()
		method := c.Request.Method
		// match1, err := regexp.MatchString("/logout", fullPath)
		path := strings.Split(c.FullPath(), "/")
		fmt.Println(path, method)
		// for links only check if people already login or not for permission to edit will be checked
		// on the next method
		if path[1] == "users" {
			if path[len(path)-1] == "logout" {
				return
			}
			if val != path[len(path)-1] {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Message": "Authorization not met"})
				return
			}
		}

		c.Next()

	}
}
