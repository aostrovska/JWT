package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
	 "os"
	 "github.com/dgrijalva/jwt-go"
)



type logPas struct{
	Username string `login`
	Password string `password`
}

var user = logPas{"username", "password"}

var token string
var times int64

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")
}

func Handler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(204)
	}else if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		var person logPas
		err = json.Unmarshal(data, &person)
		if err != nil{
			return 
		}
		if (person.Username == user.Username)&&(person.Password == user.Password){
			token, err = CreateToken(1)
			if err != nil{
				return
			}
			w.Header().Set("Token", token)
		}else {
			io.WriteString(w, "incorrect login")
		}
	}
}

func CreateToken(userId uint64) (string, error) {
  var err error
  os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userId
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  times =  time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}

func getTokenRemainingValidity(timestamp interface{}) int {
    var expireOffset = 540
    if validity, ok := timestamp.(int64); ok {
        tm := time.Unix(int64(validity), 0)
        remainder := tm.Sub(time.Now())

        if remainder > 0 {
            return int(remainder.Seconds() + float64(expireOffset))
        }
    }
    return expireOffset
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(204)
	}else if req.Method == "GET" {
		fmt.Println("req:",(*req).Header.Get("Token"))
		if ((*req).Header.Get("Token") == token)&&(getTokenRemainingValidity(times)>0){
			io.WriteString(w, "you succcessfuly gained data")
		}else{
			io.WriteString(w, "401")
		}
	}


}


func main() {
	http.HandleFunc("/login", Handler)
	http.HandleFunc("/data", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
