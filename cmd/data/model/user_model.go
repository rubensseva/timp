package model

// User represents a user
type User struct {
	username    string
	gamesPlayed int
	avgWPM      float32
}

type UserJSON struct {
	Username    string
	GamesPlayed int
	AvgWPM      float32
}

func NewUser(username string, gamesPlayed int, avgWPM float32) User {
	return User{username, gamesPlayed, avgWPM}
}

func NewUserCopy(u User) User {
	return User{u.username, u.gamesPlayed, u.avgWPM}
}

func (u User) ToJSONobj() UserJSON {
	return UserJSON{u.username, u.gamesPlayed, u.avgWPM}
}

func (u UserJSON) ToRegularObj() User {
	return User{u.Username, u.GamesPlayed, u.AvgWPM}
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetGamesPlayed() int {
	return u.gamesPlayed
}

func (u User) GetAvgWPM() float32 {
	return u.avgWPM
}

func UsersListToJSON(users []User) []UserJSON {
	var usersJSON []UserJSON
	for _, user := range users {
		usersJSON = append(usersJSON, user.ToJSONobj())
	}
	return usersJSON
}

func UsersJSONListToRegular(usersJSON []UserJSON) []User {
	var users []User
	for _, userJSON := range usersJSON {
		users = append(users, userJSON.ToRegularObj())
	}
	return users
}
