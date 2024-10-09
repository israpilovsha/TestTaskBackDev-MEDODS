package main

import (
	"fmt"
	"log"
	"net/http"
	//jwt "github.com/dgrijalva/jwt-go"
)

/*
	func GenerateJWT() (string, error) {
		token := jwt.New(jwt.SigningMethodHS512)
	}
*/

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
	fmt.Println("My service with JWT")
}
