package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type jwToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {

	var mySigningKey = []byte("123456789xfs")

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["Username"] = "userdemo"
	claims["Email"] = "test@test.com"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	fmt.Println("Bearer", tokenString)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	decoded := context.Get(r, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func main() {
	// server
	router := mux.NewRouter()
	fmt.Println("Starting Application")
	// indirizzo autenticazione
	router.HandleFunc("/authenticaticate", CreateTokenEndpoint).Methods("POST")
	//indirizzi protetti
	router.HandleFunc("/protected", ValidateMiddleware(ProtectedEndpoint)).Methods("GET")
	// define url and port of running server
	http.ListenAndServe(":8080", router)
	fmt.Println("Server running on http://127.0.0.1:8080")

}
