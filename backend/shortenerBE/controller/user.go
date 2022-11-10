package controller

import (
	"shortenerBE/helper"
	"shortenerBE/model"
)

func Login(username string, password string) {

}

func GetUserbyID(username string) {
	model.GetUserbyID(username)
}

func Register(full_name string, username string, password string) (string, error) {
	salt := helper.GenerateHash()
	hashed_password, err := helper.EncryptPassword(salt, password)
	// err.Error() to get error message
	if err != nil {
		return "failed to adding user", err
	}
	model.AddUser(full_name, username, salt, hashed_password)
	return "success", nil

}
