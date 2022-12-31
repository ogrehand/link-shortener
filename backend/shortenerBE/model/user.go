package model

import (
	"context"
	"fmt"
	"math"
	"shortenerBE/helper"
	"strings"
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
	Password   string    `json:"password,omitempty" bson:"password,omitempty"`
	Salt       string    `json:"salt,omitempty" bson:"salt,omitempty"`
	Created_at time.Time `json:"created_at" bson:"created_at,omitempty"`
	Status     bool      `json:"status" bson:"status,omitempty"`
	Token      []token   `bson:"token,omitempty"`
}

var TokenTTL time.Duration = time.Duration(86400 * math.Pow(10, 9))

func AddUser(binder func(any) error) error {
	var userData user

	binder(&userData)
	userData.Salt = helper.GenerateSalt()
	hashed_password, err := helper.EncryptPassword(userData.Salt, userData.Password)
	userData.Password = hashed_password
	userData.Created_at = time.Now()
	userData.Status = true
	userData.Token = []token{}

	usersCollection, err := ConnectDB("user")
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), userData)
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

func PanicSample() {
	panic("a")
}

func GetPassSalt(username string) (string, string) {
	usersCollection, err := ConnectDB("user")
	if err != nil {
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

func Login(username string, password string, tokenKey string) (*mongo.UpdateResult, time.Time) {
	usersCollection, err := ConnectDB("user")
	if err != nil {
		panic(err)
	}
	user := bson.M{"_id": username}
	nowDate := time.Now()

	tokenData := token{tokenKey, nowDate}
	query := bson.M{"$push": bson.M{"token": tokenData}}
	result, err2 := usersCollection.UpdateOne(context.TODO(), user, query)
	if err2 != nil {
		panic(err2)
	}
	return result, nowDate
}

func LoginRedis(username string) (string, int) {
	rdb := ConnectRedis()
	ctx := context.TODO()
	result := ""
	token := ""
	for result != username {
		fmt.Println(result, token)
		token = helper.GenerateToken()
		err := rdb.SetNX(ctx, token, username, TokenTTL).Err()
		result = strings.Split(rdb.Get(ctx, token).String(), " ")[2]
		// fmt.Println(strings.Join(result))
		if err != nil {
			panic(err)
		}
	}

	return token, int(TokenTTL.Seconds())
}

func LogoutRedis(token string) string {
	rdb := ConnectRedis()
	ctx := context.TODO()
	err := rdb.Del(ctx, token).Err()
	if err != nil {
		panic(err)
	}
	return "success"
}

func CheckTokenRedis(token string) (string, error) {
	rdb := ConnectRedis()
	ctx := context.TODO()
	val, err := rdb.Get(ctx, token).Result()
	return val, err
}

func Logout(username string, token string) *mongo.UpdateResult {
	usersCollection, err := ConnectDB("user")
	if err != nil {
		panic(err)
	}
	user := bson.M{"_id": username}
	query := bson.M{"$pull": bson.M{"token": bson.M{"key": token}}}
	result, err2 := usersCollection.UpdateOne(context.TODO(), user, query)
	if err2 != nil {
		panic(err2)
	}
	return result
}

func GetUserbyID(username string) (user, error) {

	var userData user

	usersCollection, err := ConnectDB("user")
	usersCollection.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&userData)
	userData.Password = ""
	userData.Salt = ""

	return userData, err
}
func EditUser(binder func(any) error, id string, delete bool) error {
	var userData user
	if binder != nil {
		binder(&userData)
	}

	usersCollection, err := ConnectDB("user")
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	if delete {
		userData2 := bson.M{"status": false}
		result, err := usersCollection.UpdateOne(context.TODO(), bson.M{"_id": id},
			bson.M{"$set": userData2})

		fmt.Println(result.ModifiedCount)
		return err
	}

	result, err := usersCollection.UpdateOne(context.TODO(), bson.M{"_id": id},
		bson.M{"$set": userData})

	if err != nil {
		fmt.Println(mongo.IsDuplicateKeyError(err))
		fmt.Println(err.Error())
		return err

	}
	// display the id of the newly inserted object
	fmt.Println(result)
	fmt.Println(result.MatchedCount)
	fmt.Println(result.ModifiedCount)
	return nil
}
