package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetLink(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
	fmt.Println(client, ctx, cancel, err)

	linksCollection := client.Database("test").Collection("link")
	hasil := linksCollection.FindOne(context.TODO(), bson.M{"_id": id})
	return hasil, nil

}
func AddLink(full_name string, username string, email string,
	salt string, hashed_password string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
	fmt.Println(client, ctx, cancel, err)

	usersCollection := client.Database("test").Collection("user")
	user := bson.D{{"_id", username},
		{"full_name", full_name},
		{"email", email},
		{"salt", salt},
		{"password", hashed_password},
		{"created_at", time.Now()},
		{"status", true},
		{"token", bson.A{}}}
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		fmt.Println(mongo.IsDuplicateKeyError(err))
		fmt.Println(err.Error())
		// panic(err.Error())
		return err

	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
	return nil
}
