package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"timp/cmd/data/model"
)

func readAllUsers() []model.User {
	usersfile, _ := ioutil.ReadFile("cmd/data/json/users.json")
	var users []model.User
	_ = json.Unmarshal([]byte(usersfile), &users)
	if len(users) == 0 {
		panic("Trying to get users, but no user is created. Create a user first.")
	}
	return users
}

func readAllUsersUnsafe() []model.User {
	usersfile, _ := ioutil.ReadFile("cmd/data/json/users.json")
	var users []model.User
	_ = json.Unmarshal([]byte(usersfile), &users)
	return users
}

// GetAllUsers returns all users from json file
func GetAllUsers() []model.User {
	return readAllUsers()
}

// GetUser by username
func GetUser(username string) model.User {
	var users = readAllUsersUnsafe()
	var user model.User
	var found = false
	for _, user := range users {
		if username == user.GetUsername() {
			user = model.NewUserCopy(user)
			found = true
		}
	}
	if found == false {
		fmt.Print("Warning! couldnt find user with username: " + username + ", returning user with zero-values")
	}
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
	writefile, _ := json.MarshalIndent(users, "", " ")
	_ = ioutil.WriteFile("cmd/data/json/users.json", writefile, 0644)
	fmt.Println("create user success (hopefully)")
}
