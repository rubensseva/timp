package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"timp/cmd/data/model"
)

func readLoggedInUser() model.CurrentUser {
	currentuserfile, _ := ioutil.ReadFile("cmd/data/json/currentUser.json")
	var currentUser model.CurrentUser
	_ = json.Unmarshal([]byte(currentuserfile), &currentUser)
	if currentUser.User.GetUsername() == "" {
		panic("Tried to get currently logged in user, but no name is set")
	}
	return currentUser
}

func readLoggedInUserUnsafe() model.CurrentUser {
	currentuserfile, _ := ioutil.ReadFile("cmd/data/json/currentUser.json")
	var currentUser model.CurrentUser
	_ = json.Unmarshal([]byte(currentuserfile), &currentUser)
	return currentUser
}

// GetLoggedInUser returns the currently logged in user
func GetLoggedInUser() model.CurrentUser {
	return readLoggedInUser()
}

// LogoutUser logs out whichever user is currently logged in
func LogoutUser() {
	var tmpUser = model.NewUser("not-logged-in", 0, 0.0)
	var data = model.CurrentUser{IsLoggedIn: false, User: tmpUser}
	writefile, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("cmd/data/json/currentUser.json", writefile, 0644)
}

// LoginUser logs in a user
// checks if another user is logged in, or if this user is already logged in
func LoginUser(newActiveUser model.User) {

	var users = readAllUsersUnsafe()
	var currentUser = readLoggedInUserUnsafe()

	if currentUser.IsLoggedIn {
		fmt.Println("already logged in as: ", currentUser.User.GetUsername())
		return
	}

	var isAUser = false
	for _, user := range users {
		if user.GetUsername() == newActiveUser.GetUsername() {
			isAUser = true
		}
	}

	if !isAUser {
		fmt.Println("specified username is not a user. Is the username right? Is the user created?")
		return
	}

	fmt.Println("loging in as ", newActiveUser.GetUsername())

	var tmpUser = model.NewUserCopy(newActiveUser)
	var data = model.CurrentUser{IsLoggedIn: true, User: tmpUser}
	writefile, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("cmd/data/json/currentUser.json", writefile, 0644)
	fmt.Println("loggin succes (hopefully)")
}
