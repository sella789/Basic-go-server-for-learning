package main
import (
  "net/http"
  "encoding/json"
  "fmt"
)

var index int = 0;
var users [1000]*user;
var userChats []UserChat

// The post method of users
// Adds the user to the users list
func addUser(w http.ResponseWriter, r *http.Request) {
	user := user{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}
	u := NewUser(user.Username)

  users[index] = u
  index++
}



func updateUser(w http.ResponseWriter, r *http.Request){
		user := user{} //initialize empty user

	//Parse json request body and use it to set fields on user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}

	u := NewUserWithId(user.Id, user.Username)

	for i:=0; i < 1000; i++ {
		if(u.Id == users[i].Id){
			users[i] = u
		}
	}
}

// Users get request function
// Parses the current users to json and writes 
func getUsers(w http.ResponseWriter, r *http.Request){
	js, err := json.Marshal(users[0:index+1])
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

// Routs each request of 'users' to its method
func usersRouter(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w,r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func getMessages(w http.ResponseWriter, r *http.Request){
	js, err := json.Marshal(userChats)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func messagesRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessages(w,r)
	case http.MethodPost:

	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func main() {
	// Default user
	var myMsg []message
	myMsg = append(myMsg, (message{SenderID: 1, ReceiverID:2,Content:"hello"}))
	users[0] = NewUser("sella")
	userChats = append(userChats, UserChat{
		UserID : 1,
		SecondUserID : 2,
		Messages : myMsg,
	})

	// Http server conf
	http.HandleFunc("/users", usersRouter)
	http.HandleFunc("/messages" , messagesRouter)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
