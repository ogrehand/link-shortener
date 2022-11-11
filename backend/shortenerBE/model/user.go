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

func AddUser(full_name string, username string, email string,
	salt string, hashed_password string) {
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
		panic(err.Error())

	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
}
func Login(username string, password string, token string) {
	ctx, _ := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
	usersCollection := client.Database("test").Collection("user")
	user := bson.D{{"_id", username},
		{"password", password}}
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	// need to find how to use this method
	result, err := usersCollection.UpdateOne(context.TODO(), user, user)
	// check for errors in the insertion
	if err != nil {
		fmt.Println(mongo.IsDuplicateKeyError(err))
		fmt.Println(err.Error())
		fmt.Println(result)
		panic(err.Error())
	}
}

func GetUserbyID(username string) bson.D {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bukanroot:bukanroot@mongo:27017"))
	fmt.Println(client, ctx, cancel, err)

	var userData bson.D

	usersCollection := client.Database("test").Collection("user")
	// user := bson.D{{}}
	// user := bson.M{"_id": username}
	// insert the bson object using InsertOne()
	usersCollection.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&userData)
	// check for errors in the insertion
	// if result != nil {
	// 	panic(result)
	// }
	// display the id of the newly inserted object
	fmt.Println("Adad")
	fmt.Println(userData)
	// jsonUser, _ := json.Marshal(userData)
	return userData
}
