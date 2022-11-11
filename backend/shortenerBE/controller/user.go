package controller

import (
	"fmt"
	"net/http"
	"shortenerBE/helper"
	"shortenerBE/model"

	"github.com/gin-gonic/gin"
)

type user struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(username string, password string) {

}

func GetUserbyID(c *gin.Context) {
	c.JSON(http.StatusOK, model.GetUserbyID(c.Param("id")))
}

func Register(c *gin.Context) {
	var userData user
	if err := c.BindJSON(&userData); err != nil {
		fmt.Println(err.Error())
	}
	/**
	best way to print struct instance
	fmt.Printf("%+v\n", userData)
	fmt.Println(userData.Password)
	*/
	// res2B, _ := json.Marshal(userData)
	// fmt.Println(string(res2B))
	// controller.Register("Asdas", "adasdas", "asdasdasd")
	// fmt.Println("terserah")
	salt := helper.GenerateSalt()
	hashed_password, err := helper.EncryptPassword(salt, userData.Password)
	// err.Error() to get error message
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	model.AddUser(userData.Fullname, userData.Username, userData.Email, salt, hashed_password)
	c.JSON(http.StatusOK, gin.H{
		"message": "created successfully",
	})

}
