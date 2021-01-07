package data

import (
	"encoding/json"
	"fmt"
  "os"
	"io/ioutil"
	"github.com/rubensseva/timp/cmd/data/model"
)

func readLoggedInUser() model.CurrentUser {
	currentuserfile, fileErr := ioutil.ReadFile(os.Getenv("HOME") + "/.timp/currentUser.json")
	if fileErr != nil {
		panic(fileErr)
	}
	var currentUserJSON model.CurrentUserJSON
	JSONErr := json.Unmarshal([]byte(currentuserfile), &currentUserJSON)
	if JSONErr != nil {
		panic(JSONErr)
	}
	if currentUserJSON.User.Username == "" {
		panic("Tried to get currently logged in user, but no name is set")
	}
	return currentUserJSON.ToRegularObj()
}

func readLoggedInUserUnsafe() model.CurrentUser {
	currentuserfile, fileErr := ioutil.ReadFile(os.Getenv("HOME") + "/.timp/currentUser.json")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	var currentUserJSON model.CurrentUserJSON
	JSONErr := json.Unmarshal([]byte(currentuserfile), &currentUserJSON)
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	if currentUserJSON.User.Username == "" {
		fmt.Println("Tried to get currently logged in user, but no name is set")
	}
	return currentUserJSON.ToRegularObj()
}

// GetLoggedInUser returns the currently logged in user
func GetLoggedInUser() model.CurrentUser {
	return readLoggedInUserUnsafe()
}

// LogoutUser logs out whichever user is currently logged in
func LogoutUser() {
	var tmpUser = model.NewUser("not-logged-in", 0, 0.0)
	var data = model.NewCurrentUser(false, tmpUser)
	writefile, JSONErr := json.MarshalIndent(data.ToJSONobj(), "", " ")
	if JSONErr != nil {
		panic(JSONErr)
	}
	fileErr := ioutil.WriteFile(os.Getenv("HOME") + "/.timp/currentUser.json", writefile, 0644)
	if fileErr != nil {
		panic(fileErr)
	}

}

// LoginUser logs in a user
// checks if another user is logged in, or if this user is already logged in
func LoginUser(newActiveUser model.User) {
	fmt.Println("Attemting to login user: " + newActiveUser.GetUsername())

	var users = readAllUsersUnsafe()
	var currentUser = readLoggedInUserUnsafe()

	if currentUser.GetIsLoggedIn() {
		fmt.Println("already logged in as: ", currentUser.GetUser().GetUsername())
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
	var data = model.NewCurrentUser(true, tmpUser)
	fmt.Println(string(data.GetUser().GetUsername()))
	writefile, JSONErr := json.MarshalIndent(data.ToJSONobj(), "", " ")
	if JSONErr != nil {
		panic(JSONErr)
	}
	fileErr := ioutil.WriteFile(os.Getenv("HOME") + "/.timp/currentUser.json", writefile, 0644)
	if fileErr != nil {
		panic(fileErr)
	}
	fmt.Println("loggin succes (hopefully)")
}
