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
	salt string, hashed_password string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GenerateURI()))
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
func Login(username string, password string, token string) {
	client, err := ConnectDB()
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

type user struct {
	Fullname   string    `json:"fullname" bson:"full_name,omitempty"`
	Username   string    `bson:"_id,omitempty" json:"username"`
	Email      string    `json:"email" bson:"email,omitempty"`
	Password   string    `json:"password" bson:"password,omitempty"`
	Salt       string    `json:"salt" bson:"salt,omitempty"`
	Created_at time.Time `json:"created_at" bson:"created_at,omitempty"`
	Status     bool      `json:"status" bson:"status,omitempty"`
}

func GetUserbyID(username string) user {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GenerateURI()))
	fmt.Println(client, ctx, cancel, err)

	var userData bson.D

	usersCollection := client.Database("test").Collection("user")
	usersCollection.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&userData)

	jsonUser, _ := bson.Marshal(userData)
	var userJson user
	err = bson.Unmarshal(jsonUser, &userJson)
	return userJson
}
