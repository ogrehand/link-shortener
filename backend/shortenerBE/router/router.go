package router

import(
	"io/ioutil"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteV1(router *gin.Engine){
	v1 := router.Group("/v1")
  {
    v1.GET("/login", func(c *gin.Context) {
		fmt.Println("terserah")
	})
    v1.GET("/logout", func(c *gin.Context) {
		fmt.Println("terserah 2")
	})
    v1.GET("/logup", func(c *gin.Context) {
		fmt.Println("terserah 3")
	})
  }
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