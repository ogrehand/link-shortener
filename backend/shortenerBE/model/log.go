package model

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

/**
to save error and other log
ok to be fail

next will add option to add to kafka broker
*/
type request struct {
	Header map[string][]string `bson:"header"`
	Body   string              `bson:"body"`
}
type response struct {
	Header map[string][]string `bson:"header"`
	Body   string              `bson:"body"`
}

type log struct {
	ClientAddr string    `bson:"client_addr"`
	ReqURI     string    `bson:"req_uri"`
	Method     string    `bson:"method"`
	Request    request   `bson:"request"`
	Response   response  `bson:"response"`
	Start      time.Time `bson:"start"`
	End        time.Time `bson:"end"`
	Log        string    `bson:"log"`
	Error      []string  `bson:"error"`
}
type E struct {
	Events string
}

func SaveLog(c *gin.Context, startTime time.Time, endTime time.Time,
	logString string, err []string) {
	req := c.Request
	logsCollection, _ := ConnectDB("logs")
	var bodyData E
	// var resBody E
	c.Bind(bodyData)

	logData := log{ClientAddr: req.RemoteAddr,
		ReqURI:   req.URL.String(),
		Method:   req.Method,
		Request:  request{Header: req.Header, Body: bodyData.Events},
		Response: response{Header: nil, Body: bodyData.Events},
		Start:    startTime,
		End:      endTime,
		Log:      logString,
		Error:    err,
	}
	// if &req.Response != nil {
	// 	var resBody E

	// 	logData.Response = response{Header: nil, Body: resBody.Events}
	// 	if &req.Response.Header != nil {
	// 		logData.Response.Header = req.Response.Header
	// 	}
	// }
	logsCollection.InsertOne(context.TODO(), logData)
}

// func SaveLog(errmsg string, incidentTime time.Time) {
// 	logsCollection, _ := ConnectDB("logs")

// 	logsCollection.InsertOne(context.TODO(),
// 		bson.D{{"message", errmsg}, {"time", incidentTime}})
// }
