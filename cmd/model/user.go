package model

type User struct {
	Username  string
	Highscore int
}

type CurrentUser struct {
	IsLoggedIn string
	Username   string
}
