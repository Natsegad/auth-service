package service

import (
	"auth/auth-back/internal/service/reguser"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadUsers() reguser.UsersJson {
	usersJson := reguser.UsersJson{}

	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		return usersJson
	}

	err = json.Unmarshal(file, &usersJson)
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
		return usersJson
	}

	return usersJson
}
