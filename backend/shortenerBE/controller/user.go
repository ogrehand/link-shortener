package controller

import (
	"fmt"
	"shortenerBE/helper"
	"shortenerBE/model"
)

func Login(username string, password string) {

}
func Register(full_name string, username string, password string) string {
	salt := helper.GenerateHash()
	hashed_password, err := helper.EncryptPassword(salt, password)
	// err.Error() to get error message
	fmt.Println(hashed_password)
	fmt.Println("bedain")
	if err != nil {
		return "failed to adding user"
	}
	model.AddUser(full_name, username, salt, hashed_password)
	return "success"
	// fmt.Println(hasil_pass)

}
