package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"shortenerBE/controller"

	"github.com/gin-gonic/gin"
	// "net/http"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

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
			users.POST("/login", func(c *gin.Context) {
				fmt.Println("terserah")
			})
			users.POST("/logout", func(c *gin.Context) {
				fmt.Println("terserah 2")
			})
			users.GET("/:id", func(c *gin.Context) {
				controller.GetUserbyID(c.Param("id"))
			})
			users.PUT("/", func(c *gin.Context) {
				var userData user
				if err := c.BindJSON(&userData); err != nil {
					fmt.Println(err.Error())
				}
				controller.Register(userData.Fullname,
					userData.Username,
					userData.Password)
				/**
				best way to print struct instance
				fmt.Printf("%+v\n", userData)
				fmt.Println(userData.Password)
				*/
				// res2B, _ := json.Marshal(userData)
				// fmt.Println(string(res2B))
				// controller.Register("Asdas", "adasdas", "asdasdasd")
				// fmt.Println("terserah")
			})
			users.POST("/:id", func(c *gin.Context) {
				fmt.Println("terserah 3")
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
