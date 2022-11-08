package model

import (
	"os"
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func Getdata(){
	ctx, cancel := context.WithTimeout(context.Background(),
			30*time.Second)

		// mongo.Connect return mongo.Client method
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
		fmt.Println(client, ctx, cancel, err, os.Getenv("FOO"))

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
}