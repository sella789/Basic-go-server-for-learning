package main

// The id counter
var currentId int = 0

// The user struct
type user struct{
	Id int
	Username string
}

func NewUser(username string) *user{
	u := new(user)
	u.Id = currentId
	u.Username = username

	return u
}

func NewUserWithId(id int,username string) *user{
	u := new(user)
	u.Id = id
	u.Username = username

	return u
}