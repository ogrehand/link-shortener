package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"

	"context"
	"fmt"
	"log"
	"time"
	"io/ioutil"
	// "bytes"

	"github.com/gin-gonic/gin"
	"shortenerBE/router"
	"shortenerBE/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-contrib/static"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// func proxy(c *gin.Context) {

// 	remote, err := url.Parse("http://link-shortener-frontend-1:8080")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// proxy := httputil.NewSingleHostReverseProxy(remote)

// 	director := func(req *http.Request) {
// 		fmt.Println("disini   ")
// 		fmt.Println(c.Request.Header)

// 		fmt.Println("    disini")
// 		// req.Header = c.Request.Header
// 		// req.Host = remote.Host
// 		req.URL.Scheme = remote.Scheme
// 		req.URL.Host = remote.Host
// 		req.URL.Path = remote.Path
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

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

	// r.GET("/123", proxy)
	// r.GET("/js/*any", proxy)
	

	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.RouteV1(r)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(helper.GenerateHash())
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/testdb", func(c *gin.Context) {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://bukanroot:bukanroot@localhost:27017"))
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		defer client.Disconnect(ctx)
	})
	r.GET("/routetest", func(c *gin.Context) {
		router.Router()
	})

	r.GET("/dbtest", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(),
			30*time.Second)

		// mongo.Connect return mongo.Client method
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
		fmt.Println(client, ctx, cancel, err)

		usersCollection := client.Database("test").Collection("sample2")
		user := bson.D{{"fullName", "User 1"}, {"age", 30}}
		// insert the bson object using InsertOne()
		result, err := usersCollection.InsertOne(context.TODO(), user)
		// check for errors in the insertion
		if err != nil {
			panic(err)
		}
		// display the id of the newly inserted object
		fmt.Println(result.InsertedID)
		c.JSON(http.StatusOK, gin.H{
			"message":  client,
			"message2": ctx,
		})

	})

	r.NoRoute(func(ctx *gin.Context) {
        file, _ := ioutil.ReadFile("./views/index.html")
        // etag := fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
        // ctx.Header("ETag", etag)
        ctx.Header("Cache-Control", "no-cache")

        // if match := ctx.GetHeader("If-None-Match"); match != "" {
        //     if strings.Contains(match, etag) {
        //         ctx.Status(http.StatusNotModified)

        //         //這裡若沒 return 的話，會執行到 ctx.Data
        //         return
        //     }
        // }

        ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
    })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
