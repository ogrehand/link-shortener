package model

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func GenerateURI() string {
	return "mongodb://" + os.Getenv("user_db") + ":" + os.Getenv("pass_db") +
		"@" + os.Getenv("hostname") + ":" + os.Getenv("port")
}

func ConnectDB(collection string) (*mongo.Collection, error) {
	ctx, _ := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GenerateURI()))
	if err != nil {
		return nil, err
	} else {
		return client.Database(os.Getenv("database")).Collection(collection), nil
	}
}

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host") + ":" + os.Getenv("redis_port"), // use default Addr
		Password: "",                                                      // no password set
		DB:       0,                                                       // use default DB
	})
	// fmt.Println("bisa jalan")
	// ctx := context.TODO()
	// fmt.Println("bisa jalan 1")

	// pong, err := rdb.Ping(ctx).Result()
	// fmt.Println("bisa jalan 2")
	// fmt.Println(pong, err)
	return rdb
}

func Getdata() {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GenerateURI()))
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
