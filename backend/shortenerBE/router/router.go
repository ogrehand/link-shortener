package router

import (
	"shortenerBE/controller"

	"github.com/gin-gonic/gin"
)

func RouteV1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
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
	}
	router.GET(":id", controller.Redirect)
}
