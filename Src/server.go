package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var index int
var users []*user
var userChats []UserChat

// The post method of users
// Adds the user to the users list
func addUser(w http.ResponseWriter, r *http.Request) {
	user := user{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	u := NewUser(user.Username)

	users = append(users, u)
	index++
}

/**
 * * updateUser function
 * * the function updates an existing user
 * @param w the response writer a basic param in http requests
 * @param r the request a basic param in http requests
 */
func updateUser(w http.ResponseWriter, r *http.Request) {
	user := user{} //initialize empty user

	//Parse json request body and use it to set fields on user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	u := NewUserWithID(user.ID, user.Username)

	for i := range users {
		if u.ID == users[i].ID {
			users[i] = u
		}
	}
}

// Users get request function
// Parses the current users to json and writes
func getUsers(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //! Parsing error
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

// Routs each request of 'users' to its method
func UsersRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(userChats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

/**
 * * messageRouter function
 * * the function routes the request according
 * * to the method in the request
 * @param w the response writer a basic param in http requests
 * @param r the request a basic param in http requests
 */
func messagesRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessages(w, r)
	case http.MethodPost:

	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func LoadFile(w http.ResponseWriter, r *http.Request) {

}

/**
 * * The main function of the program
 * * The function inits the server
 */
func main() {

	// Test data for get request
	var myMsg []message
	myMsg = append(myMsg, (message{SenderID: 1, ReceiverID: 2, Content: "hello"}))
	users = append(users, NewUser("stella"))
	userChats = append(userChats, UserChat{
		UserID:       1,
		SecondUserID: 2,
		Messages:     myMsg,
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// Http server conf

	http.HandleFunc("/users", UsersRouter)
	http.HandleFunc("/messages", messagesRouter)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
