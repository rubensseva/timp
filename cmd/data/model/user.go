package model

// User represents a user
type User struct {
	username    string
	gamesPlayed int
	avgWPM      float32
}

func NewUser(username string, gamesPlayed int, avgWPM float32) User {
	return User{username, gamesPlayed, avgWPM}
}

func NewUserCopy(u User) User {
	return User{u.username, u.gamesPlayed, u.avgWPM}
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
