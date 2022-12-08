package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/**
to save error and other log
ok to be fail

next will add option to add to kafka broker
*/
func SaveLog(errmsg string, incidentTime time.Time) {
	logsCollection, _ := ConnectDB("logs")

	logsCollection.InsertOne(context.TODO(),
		bson.D{{"message", errmsg}, {"time", incidentTime}})
}
