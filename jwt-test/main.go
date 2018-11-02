package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
		json.NewEncoder(w).Encode(error)
	}
	json.NewEncoder(w).Encode(jwToken{Token: tokenString})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// server
	router := mux.NewRouter()
	fmt.Println("Starting Application")
	// indirizzo autenticazione
	router.HandleFunc("/authenticaticate", CreateTokenEndpoint).Methods("POST")
	//indirizzi protetti
	router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	// define url and port of running server
	http.ListenAndServe(":8080", router)
	fmt.Println("Server running on http://127.0.0.1:8080")

}
