package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"shortenerBE/helper"
	"shortenerBE/router"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://link-shortener-frontend-1:8080")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		fmt.Println(c.Request.Header)
		fmt.Println(remote.Host)
		fmt.Println(remote.Scheme)
		fmt.Println(remote.Host)
		fmt.Println(remote.Path)
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.RouteV1(r)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(helper.GenerateSalt())

		fmt.Println(os.Getenv("user_db"))
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		file, _ := ioutil.ReadFile("./views/index.html")
		// etag := fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
		// ctx.Header("ETag", etag)
		ctx.Header("Cache-Control", "no-cache")

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
