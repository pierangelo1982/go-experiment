package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type jwToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

var signingKey = []byte("signing-key")

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, User{})
	ss, err := token.SignedString(signingKey)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		log.Printf("err: %+v\n", err)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(ss))
	log.Printf("issued token: %v\n", ss)
	return
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {}

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
