package model

// CurrentUser represents the currently logged in user
type CurrentUser struct {
	isLoggedIn bool
	user       User
}

func NewCurrentUser(isLoggedIn bool, user User) CurrentUser {
  return CurrentUser{isLoggedIn, user}
}

func NewCurrentUserCopy(c CurrentUser) CurrentUser {
  return CurrentUser{ c.isLoggedIn, c.user }
}

func (c CurrentUser) GetIsLoggedIn() bool {
  return c.isLoggedIn
}

func (c CurrentUser) GetUser() User {
  return c.user
}
