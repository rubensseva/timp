package model

// CurrentUser represents the currently logged in user
type CurrentUser struct {
	isLoggedIn bool
	user       User
}

type currentUserJSON struct {
  IsLoggedIn bool
  User      UserJSON
}

func NewCurrentUser(isLoggedIn bool, user User) CurrentUser {
  return CurrentUser{isLoggedIn, user}
}

func NewCurrentUserCopy(c CurrentUser) CurrentUser {
  return CurrentUser{ c.isLoggedIn, c.user }
}

func (u CurrentUser) ToJSONobj() currentUserJSON {
  return currentUserJSON{u.isLoggedIn, u.user.ToJSONobj()}
}

func (c CurrentUser) GetIsLoggedIn() bool {
  return c.isLoggedIn
}

func (c CurrentUser) GetUser() User {
  return c.user
}

