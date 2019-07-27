package data

import (
	"encoding/json"
	"io/ioutil"
	"timp/cmd/model"
)

func readLoggedInUser() model.CurrentUser {
	currentuserfile, _ := ioutil.ReadFile("cmd/data/json/currentUser.json")
	var currentUser model.CurrentUser
	_ = json.Unmarshal([]byte(currentuserfile), &currentUser)
	return currentUser
}

// GetLoggedInUser returns the currently logged in user
func GetLoggedInUser() model.CurrentUser {
	return readLoggedInUser()
}
