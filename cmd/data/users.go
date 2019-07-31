package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"timp/cmd/data/model"
)

func readAllUsers() []model.User {
	usersfile, fileErr := ioutil.ReadFile("cmd/data/json/users.json")
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
	usersfile, fileErr := ioutil.ReadFile("cmd/data/json/users.json")
  if fileErr != nil {
    panic(fileErr)
  }
	var usersJSON []model.UserJSON
  JSONErr := json.Unmarshal([]byte(usersfile), &usersJSON)
  if JSONErr != nil {
    panic(JSONErr)
  }
	return model.UsersJSONListToRegular(usersJSON)
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
		fmt.Print("Warning! couldnt find user with username: " + username + " , returning user with zero-values")
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

	writefile, JSONErr := json.MarshalIndent(model.UsersListToJSON(users), "", " ")
  if JSONErr != nil {
    panic(JSONErr)
  }
  fileErr := ioutil.WriteFile("cmd/data/json/users.json", writefile, 0644)
  if fileErr != nil {
    panic(fileErr)
  }
	fmt.Println("create user success (hopefully)")
}
