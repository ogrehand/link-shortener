package main

import (
	"net/http"

	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"shortenerBE/router"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	r := gin.Default()
	router.RouteV1(r)
	r.GET("/ping", func(c *gin.Context) {
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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
