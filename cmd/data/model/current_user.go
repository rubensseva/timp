package model

// CurrentUser represents the currently logged in user
type CurrentUser struct {
	IsLoggedIn bool
	User       User
}
