package controller

import (
	"io/ioutil"
	"net/http"
	"shortenerBE/model"

	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	id := c.Param("id")
	linkObj := model.GetLink(id)
	c.JSON(http.StatusOK, linkObj)
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

func DeleteLink(c *gin.Context) {
	id := c.Param("id")
	model.EditLink(nil, id, true)
}

func UpdateLink(c *gin.Context) {
	id := c.Param("id")
	model.EditLink(c.BindJSON, id, false)
}

func AddLink(c *gin.Context) {
	errDb := model.AddLink(c.BindJSON)
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
