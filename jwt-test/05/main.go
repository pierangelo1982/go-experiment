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
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Exception struct {
	Message string `json:"message"`
}

var mySigningKey = []byte("sopralapancalacapracantasottolapancalacapracrepa")

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Benvenuto nella nostra Home Page"))
}

func login(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["Username"] = "userdemo"
	claims["Password"] = "password123"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, _ := token.SignedString(mySigningKey)

	w.Write([]byte(tokenString))
}

func areaRiservata(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Benvenuto nella tua area riservata"))
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(mySigningKey), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(w, r)
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
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/area", ValidateMiddleware(areaRiservata)).Methods("GET")
	http.ListenAndServe(":8080", router)
	fmt.Println("il server è attivo all'indirizzo http://127.0.0.1:8080")
}
