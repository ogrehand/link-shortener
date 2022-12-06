package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shortenerBE/helper"
	"shortenerBE/model"

	"github.com/gin-gonic/gin"
)

type collaborator struct {
	CollaboratorId string `json:"collaborator"`
	Role           int    `json:"role"`
}
type link struct {
	Id           string `json:"shorturl"`
	RealLink     string `json:"realLink"`
	Collaborator []*collaborator
}

func Redirect(c *gin.Context) {
	id := c.Param("id")
	linkObj := model.GetLink(id)
	if linkObj.Status {
		c.Redirect(http.StatusMovedPermanently, linkObj.RealLink)
	} else {
		file, _ := ioutil.ReadFile("./views/index.html")
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	}
}

func RandomRoute(c *gin.Context) {
	id := c.Param("id")
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/"+id)
}

func DeleteLink(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func AddLink(c *gin.Context) {
	var userData user
	if err := c.BindJSON(&userData); err != nil {
		fmt.Println(err.Error())
	}
	salt := helper.GenerateSalt()
	hashed_password, err := helper.EncryptPassword(salt, userData.Password)
	// err.Error() to get error message
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	errDb := model.AddUser(userData.Fullname, userData.Username, userData.Email, salt, hashed_password)
	if errDb != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": errDb.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "created successfully",
		})
	}
}
