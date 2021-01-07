package data

import (
	"encoding/json"
	"fmt"
  "os"
	"io/ioutil"
	"github.com/rubensseva/timp/cmd/data/model"
)

func readAllUsers() []model.User {
	usersfile, fileErr := ioutil.ReadFile(os.Getenv("HOME") + "/.timp/users.json")
	if fileErr != nil {
		panic(fileErr)
	}
	var usersJSON []model.UserJSON
	JSONErr := json.Unmarshal([]byte(usersfile), &usersJSON)
	if JSONErr != nil {
		panic(JSONErr)
	}
	if len(usersJSON) == 0 {
		panic("Trying to get users, but no user is created. Create a user first.")
	}
	return model.UsersJSONListToRegular(usersJSON)
}

func readAllUsersUnsafe() []model.User {
	usersfile, fileErr := ioutil.ReadFile(os.Getenv("HOME") + "/.timp/users.json")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	var usersJSON []model.UserJSON
	JSONErr := json.Unmarshal([]byte(usersfile), &usersJSON)
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	if len(usersJSON) == 0 {
		fmt.Println("Trying to get users, but no user is created. Create a user first.")
	}
	return model.UsersJSONListToRegular(usersJSON)
}

// GetAllUsers returns all users from json file
func GetAllUsers() []model.User {
	return readAllUsersUnsafe()
}

// GetUser by username
func GetUser(username string) model.User {
	var users = readAllUsersUnsafe()
	var user model.User
	var found = false
	fmt.Println("")
	fmt.Println("")
	for _, u := range users {
		fmt.Println(username, user.GetUsername())
		if username == u.GetUsername() {
			user = model.NewUserCopy(u)
			found = true
		}
	}
	if found == false {
		fmt.Println("Warning! couldnt find user with username: " + username + " , returning user with zero-values")
	} else {
		if user.GetUsername() == "" {
			fmt.Println("Warning! User was found, but username string is empty")
		}
		fmt.Println("In GetUser, found user with username: " + username + " , returning user")
	}
	fmt.Println("In GetUser, found user with username: " + user.GetUsername() + " , returning user")
	return user
}

// AddUser gets all users and adds a user if the username
// isnt already taken
func AddUser(newUser model.User) {

	var users = readAllUsersUnsafe()

	var isAUser = false
	for _, user := range users {
		if user.GetUsername() == newUser.GetUsername() {
			isAUser = true
		}
	}

	if isAUser {
		fmt.Println("specified username " + newUser.GetUsername() + " is already a user")
		return
	}

	fmt.Println("creating user ", newUser.GetUsername())

	users = append(users, newUser)

	writefile, JSONErr := json.MarshalIndent(model.UsersListToJSON(users), "", " ")
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	fileErr := ioutil.WriteFile(os.Getenv("HOME") + "/.timp/users.json", writefile, 0644)
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	fmt.Println("create user success (hopefully)")
}
