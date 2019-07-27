package model

// CurrentUser represents the currently logged in user
type CurrentUser struct {
	IsLoggedIn string
	Username   string
}

func getIsLoggedIn(u CurrentUser) string {
	return u.IsLoggedIn
}
