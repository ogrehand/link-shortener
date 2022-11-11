package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"shortenerBE/controller"

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
			users.POST("/login", func(c *gin.Context) {
				fmt.Println("terserah")
			})
			users.POST("/logout", func(c *gin.Context) {
				fmt.Println("terserah 2")
			})

		}

		links := v1.Group("/links")
		{
			links.PUT("/", func(c *gin.Context) {
				fmt.Println("terserah 2")
			})
			links.GET("/", func(c *gin.Context) {
				fmt.Println("terserah 2")
			})
			links.POST("/:id", func(c *gin.Context) {
				fmt.Println("links 1")
			})
			links.GET("/:id", func(c *gin.Context) {
				fmt.Println("terserah 3")
			})
		}
	}
	// router.GET(":id",func(c *gin.Context){
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// })
}

func Terserah(apa string) {
	fmt.Println(apa)
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
