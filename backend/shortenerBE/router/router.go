package router

import(
	"io/ioutil"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func RouteV1(router *gin.Engine){
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
				fmt.Println("terserah 3")
			})
			users.PUT("/:id", func(c *gin.Context) {
				fmt.Println("terserah 3")
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

func Terserah(apa string){
	fmt.Println(apa)
}

func Router(){
	files, err := ioutil.ReadDir("./")
	if err != nil {
        log.Fatal(err)
    }

	for _, f := range files {
		fmt.Println(f.Name())
	}	
}