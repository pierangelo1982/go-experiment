package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const username string = "xxxxxx"
const password string = "xxxxx"
const loginUrl = "http://api.xxxxx"

var myToken string

type credenziali struct {
	Token string
}

func main() {
	fmt.Println(username)
	fmt.Println(password)
	myToken := getToken(loginUrl, username, password)
	fmt.Println(myToken)
	//b, err := json.Marshal("http://api.fintyreclub.it/gommista/token")
}

func getToken(url string, username string, password string) string {
	var token string
	client := http.Client{}
	var jsonprep string = `{"username":"` + username + `","password":"` + password + `"}`
	var jsonStr = []byte(jsonprep)
	req, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("body=", string(body))
		data := credenziali{}
		_ = json.Unmarshal([]byte(body), &data)
		token = data.Token

	}
	fmt.Println("token:", token)
	return token
}
