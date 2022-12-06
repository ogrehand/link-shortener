package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"shortenerBE/controller"
	"shortenerBE/model"

	"github.com/gin-gonic/gin"
)

type user struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RouteV1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/:id", controller.GetUserbyID)
			users.PUT("/", controller.Register)
			users.POST("/:id", func(c *gin.Context) {
				fmt.Println("terserah 3")
			})
			users.POST("/login", controller.Login)
			users.POST("/logout", controller.Logout)

		}

		links := v1.Group("/links")
		{
			links.PUT("/", func(c *gin.Context) {
				var collaborators []model.Collaborator
				model.AddLink("asd", "https://google.com", "tralala23", true, collaborators)
			})
			links.GET("/", func(c *gin.Context) {
				fmt.Println("terserah 2")
			})
			links.POST("/:id", func(c *gin.Context) {
				fmt.Println("links 1")
			})
			links.GET("/:id", func(c *gin.Context) {
				fmt.Printf("%+v\n", model.GetLink(c.Param("id")))
			})
		}
	}
	// router.GET("/dbdicoba", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": model.GenerateURI(),
	// 	})
	// })
	router.GET(":id", controller.Redirect)
}

func Router() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
