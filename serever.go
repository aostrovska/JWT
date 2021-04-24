package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)



type logPas struct{
	Username string `login`
	Password string `password`
}

var user = logPas{
  Username: "username",
  Password: "password",
}
var token string

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Handler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	data, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	var person logPas
	err = json.Unmarshal(data, &person)
	if err != nil{
		return 
	}
	fmt.Println("handler1")
	if (person.Username == user.Username)&&(person.Password == user.Password){
		token = "12345"
	}else {
		io.WriteString(w, "incorrect login")
	}
	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	if (token == "12345"){
		io.WriteString(w, "you succcessfuly gained data")
	}else{
		io.WriteString(w, "you unsucccessfuly gained data")
	}


}


func main() {
	http.HandleFunc("/login", Handler)
	http.HandleFunc("/data", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
