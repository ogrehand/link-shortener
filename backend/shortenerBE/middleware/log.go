package middleware

import (
	"fmt"
	"shortenerBE/model"
	"time"

	"github.com/gin-gonic/gin"
)

//https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// before request

		c.Next()
		var listErr []string
		for _, ginErr := range c.Errors {
			listErr = append(listErr, ginErr.Error())
		}
		fmt.Println(c.Get("key1"))
		model.SaveLog(c, startTime, time.Now(), "bebas", listErr)
	}
}
