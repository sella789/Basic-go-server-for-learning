package main

// The id counter
var currentID int

// The user struct
type user struct {
	ID       int
	Username string
	passwrod string
}

/**
* NewUser
* * The function creates a new user
* * With a generated id
* @param username the user's username
**/
func NewUser(username string) *user {
	u := new(user)
	u.ID = currentID
	u.Username = username

	return u
}

/**
* NewUserWithId
* * The function creates a new user
* * With a given id
* @param username the user's username
* @param id the users id
**/
func NewUserWithID(id int, username string) *user {
	u := new(user)
	u.ID = id
	u.Username = username

	return u
}
