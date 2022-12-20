package router

import (
	"fmt"
	"reflect"
	"shortenerBE/controller"
	"shortenerBE/model"

	"github.com/gin-gonic/gin"
)

func RouteV1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ipfinder", func(c *gin.Context) {
			fmt.Println(reflect.TypeOf(c.Request))
			fmt.Println(reflect.TypeOf(c.Request))
			fmt.Println(reflect.TypeOf(c.Request.Response))

			fmt.Println(c.Request)
			fmt.Println(c.Request.Response)

			fmt.Println(c.Request.Header["X-Forwarded-For"])
			fmt.Println(c.Request.Header["x-forwarded-for"])
			fmt.Println(c.Request.Header["X-FORWARDED-FOR"])
			fmt.Println(c.ClientIP())
			fmt.Println(c.Request.RemoteAddr)

			for name, values := range c.Request.Header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Println(name, value)
				}
			}
			// model.SaveLog(c, time.Now(), time.Now(),
			// 	"asdasd\nsadada", errors.New("tidak aadda"))
		})
		users := v1.Group("/users")
		{
			users.GET("/:id", controller.GetUserbyID)
			users.PUT("/", controller.Register)
			users.POST("/:id", controller.EditUser)
			users.DELETE("/:id", controller.DeleteUser)
			users.POST("/login", controller.Login)
			users.POST("/logout", controller.Logout)

		}

		links := v1.Group("/links")
		{
			links.PUT("/", controller.AddLink)
			links.POST("/:id", controller.UpdateLink)
			links.GET("/:id", controller.GetLink)
			links.DELETE("/:id", controller.DeleteLink)
		}
		dbs := v1.Group("/dbs")
		{
			dbs.POST("/:ping", func(c *gin.Context) {
				model.ConnectRedis()
			})
			dbs.GET("/", func(c *gin.Context) {
				fmt.Println("router bisa")
				model.ConnectRedis()
				fmt.Println("router bisa")
			})
		}
	}
	router.GET(":id", controller.Redirect)
}
