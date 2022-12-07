package model

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Collaborator struct {
	CollaboratorId string `json:"collaborator_id" bson:"collab_id,omitempty"`
	Role           string `json:"role" bson:"role,omitempty"`
}

type Link struct {
	ShortLink    string         `json:"short_link" bson:"_id,omitempty"`
	RealLink     string         `json:"real_link" bson:"real_link,omitempty"`
	Author       string         `json:"author" bson:"author,omitempty"`
	Status       bool           `json:"status" bson:"status,omitempty"`
	Collaborator []Collaborator `json:"collaborators" bson:"collaborators,omitempty"`
}

func GetLink(id string) *Link {

	linksCollection, err := ConnectDB("link")
	if err != nil {
		fmt.Println(err.Error())
	}
	var linkObj Link

	linksCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&linkObj)
	fmt.Println(linkObj.RealLink)
	return &linkObj

}

func EditLink(binder func(any) error, shortlink string, delete bool) error {
	var linkObj Link
	if binder != nil {
		binder(&linkObj)
	}

	linksCollection, err := ConnectDB("link")
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	if delete {
		linkObj2 := bson.M{"status": false}
		result, err := linksCollection.UpdateOne(context.TODO(), bson.M{"_id": shortlink},
			bson.M{"$set": linkObj2})

		fmt.Println(result.ModifiedCount)
		return err
	}

	result, err := linksCollection.UpdateOne(context.TODO(), bson.M{"_id": shortlink},
		bson.M{"$set": linkObj})

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

func AddLink(binder func(any) error) error {
	var linkObj Link
	binder(&linkObj)
	linkObj.Author = "tralal2a3"
	linkObj.Status = true
	linkObj.Collaborator = []Collaborator{}

	if linkObj.ShortLink == "" {
		err := errors.New("short link can't be empty")
		return err
	}
	if linkObj.RealLink == "" {
		err := errors.New("real link can't be empty")
		return err
	}
	linksCollection, err := ConnectDB("link")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	result, err := linksCollection.InsertOne(context.TODO(), linkObj)

	if err != nil {
		fmt.Println(mongo.IsDuplicateKeyError(err))
		fmt.Println(err.Error())
		return err

	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
	return nil
}
