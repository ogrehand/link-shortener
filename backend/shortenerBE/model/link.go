package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Collaborator struct {
	CollaboratorId string `bson:"collab_id"`
	Role           string `bson:"role"`
}

type Link struct {
	ShortLink    string         `bson:"_id"`
	RealLink     string         `bson:"real_link"`
	Author       string         `bson:"author"`
	Status       bool           `bson:"status"`
	Collaborator []Collaborator `bson:"collaborators"`
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

func AddLink(shortlink string, reallink string, author string,
	status bool, collaborators []Collaborator) error {

	linksCollection, err := ConnectDB("link")
	if err != nil {
		fmt.Println(err.Error())
		return err

	}

	linkObj := Link{shortlink, reallink, author, status, collaborators}
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
