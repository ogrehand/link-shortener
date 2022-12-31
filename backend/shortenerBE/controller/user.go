package controller

import (
	"fmt"
	"net/http"
	"shortenerBE/helper"
	"shortenerBE/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username, existu := c.GetPostForm("username")
	salt, hash := model.GetPassSalt(username)
	password, exist := c.GetPostForm("password")

	if exist && existu && helper.CheckHash(salt, password, hash) {
		result, timeData := model.LoginRedis(username)

		c.JSON(http.StatusOK, gin.H{
			"token": result,
			"TTL":   timeData,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password wrong",
		})
	}

	// mongo token
	// if exist && existu {
	// 	helper.CheckHash(salt, password, hash)
	// 	token := helper.GenerateToken()
	// 	result, timeData := model.Login(username, hash, token)
	// 	if result.ModifiedCount == 1 {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"token":      token,
	// 			"created_at": timeData,
	// 		})
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "username or password empty",
	// 	})
	// }

}

func Logout(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")
	result := model.LogoutRedis(token[1])
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func GetUserbyID(c *gin.Context) {
	result, err := model.GetUserbyID(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, result)
}

func Register(c *gin.Context) {
	/**
	best way to print struct instance
	fmt.Printf("%+v\n", userData)
	fmt.Println(userData.Password)
	*/
	// res2B, _ := json.Marshal(userData)
	// fmt.Println(string(res2B))
	// controller.Register("Asdas", "adasdas", "asdasdasd")
	// fmt.Println("terserah")
	// err.Error() to get error message
	c.Set("key1", "you know who")
	errDb := model.AddUser(c.BindJSON)
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

func EditUser(c *gin.Context) {
	err := model.EditUser(c.BindJSON, c.Param("id"), false)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Edt successfully"})
	}
}

func DeleteUser(c *gin.Context) {
	err := model.EditUser(nil, c.Param("id"), true)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Account Deactivated"})
	}
}
