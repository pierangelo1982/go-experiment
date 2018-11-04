package main

import (
	"fmt"
	"net/http"
	"time"

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

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {

	var mySigningKey = []byte("secret")

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
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
