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

type token struct {
	Key        string    `bson:"key,omitempty"`
	Created_at time.Time `bson:"created_at,omitempty"`
}

type user struct {
	Fullname   string    `json:"fullname" bson:"full_name,omitempty"`
	Username   string    `json:"username" bson:"_id,omitempty"`
	Email      string    `json:"email" bson:"email,omitempty"`
	Password   string    `json:"password" bson:"password,omitempty"`
	Salt       string    `json:"salt" bson:"salt,omitempty"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Status     bool      `json:"status" bson:"status"`
	Token      []token   `bson:"token,omitempty"`
}

func AddUser(full_name string, username string, email string,
	salt string, hashed_password string) error {

	usersCollection, err := ConnectDB("user")
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

func GetPassSalt(username string) (string, string) {
	usersCollection, err := ConnectDB("user")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	projection := bson.D{{"_id", 0},
		{"full_name", 0},
		{"email", 0},
		{"created_at", 0},
		{"status", 0},
		{"token", 0}}

	result := usersCollection.FindOne(context.TODO(), bson.D{{"_id", username}},
		options.FindOne().SetProjection(projection))

	var saltData user
	result.Decode(&saltData)
	return saltData.Salt, saltData.Password
}

func Login(username string, password string, token string) *mongo.UpdateResult {
	usersCollection, err := ConnectDB("user")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	user := bson.M{"_id": username}
	query := bson.M{"$push": bson.M{"token": bson.D{{"key", token},
		{"created_at", time.Now()}}}}
	result, err2 := usersCollection.UpdateOne(context.TODO(), user, query)
	fmt.Println(result)
	if err2 != nil {
		panic(err2.Error())
	}
	return result
}

func Logout(username string, token string) *mongo.UpdateResult {
	usersCollection, err := ConnectDB("user")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	user := bson.M{"_id": username}
	query := bson.M{"$pull": bson.M{"token": bson.M{"key": token}}}
	result, err2 := usersCollection.UpdateOne(context.TODO(), user, query)
	fmt.Println(result)
	if err2 != nil {
		panic(err2.Error())
	}
	return result
}

func GetUserbyID(username string) (user, error) {

	var userData bson.D

	usersCollection, err := ConnectDB("user")
	usersCollection.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&userData)

	jsonUser, _ := bson.Marshal(userData)
	var userJson user
	err = bson.Unmarshal(jsonUser, &userJson)
	return userJson, err
}
