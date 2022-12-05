package controller

import (
	"errors"
	"fmt"
	"math/rand"
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

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}
func Login(c *gin.Context) {
	username, existu := c.GetPostForm("username")
	salt, hash := model.GetPassSalt(username)
	password, exist := c.GetPostForm("password")
	fmt.Println("betul ", existu, exist, " salah")
	if exist && existu {
		helper.CheckHash(salt, password, hash)
		token := helper.GenerateToken()
		result := model.Login(username, hash, token)
		if result.ModifiedCount == 1 {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password empty",
		})
	}

}

func Logout(c *gin.Context) {
	username, existu := c.GetPostForm("username")
	token := c.Request.Header["Token"]

	fmt.Println(username, existu, token)

	result := model.Logout(username, token[0])
	fmt.Println("sudah keluar sekian ", result.ModifiedCount)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetUserbyID(c *gin.Context) {
	// c.JSON(http.StatusOK, model.GetUserbyID(c.Param("id")))
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
