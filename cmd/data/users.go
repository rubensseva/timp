package data

import (
	"encoding/json"
	"io/ioutil"
	"timp/cmd/model"
)

func readAllUsers() []model.User {
	usersfile, _ := ioutil.ReadFile("cmd/data/json/users.json")
	var users []model.User
	_ = json.Unmarshal([]byte(usersfile), &users)
	return users
}

// GetAllUsers returns all users from json file
func GetAllUsers() []model.User {
	return readAllUsers()
}
