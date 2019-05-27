package main
import (
  "net/http"
  "encoding/json"
  "fmt"
)

type user struct{
	Username string
}

var index int = 0;
var users [1000]user;


func addUser(w http.ResponseWriter, r *http.Request) {
	user := user{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}
  //message = "Hello " + message
  users[index] = user
  index++
  //w.Write([]byte(message))
}

func getUsers(w http.ResponseWriter, r *http.Request){
	js, err := json.Marshal(users[0:index])
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

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

func main() {
  users[0].Username = "sella"
  http.HandleFunc("/users", usersRouter)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
