package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pierangelo1982/go-experiment/delbarbaStoreJsonInDB/model"
)

const username string = "xxxxxxxxxx"
const password string = "xxxxxxxxxx"
const loginURL = "http://api.fintyreclub.it/gommista/token"

var myToken string

type customer struct {
	ID                        int    `json:"id"`
	DataInserimento           string `json:"data_inserimento"`
	Nome                      string `json:"nome"`
	Cognome                   string `json:"cognome"`
	IDGommista                string `json:"id_gommista"`
	Indirizzo                 string `json:"indirizzo"`
	Citta                     string `json:"citta"`
	Provincia                 string `json:"provincia"`
	Cap                       string `json:"cap"`
	Telefono                  string `json:"telefono"`
	Email                     string `json:"email"`
	Cellulare                 string `json:"cellulare"`
	Note                      string `json:"note"`
	IsApp                     string `json:"is_app"`
	DataScadenzaPatente       string `json:"data_scadenza_patente"`
	DataRegistrazioneApp      string `json:"data_registrazione_app"`
	DataConfermaRegistrazione string `json:"data_conferma_registrazione_app"`
	DataUltimoLogin           string `json:"data_ultimo_login_app"`
	//Veicoli                   []string `json:"veicoli"`
}

func main() {
	fmt.Println(username)
	fmt.Println(password)
	myToken := getToken(loginURL, username, password)
	listCustomer(myToken)
	fmt.Println(myToken)
	//b, err := json.Marshal("http://api.fintyreclub.it/gommista/token")
}

func getToken(url string, username string, password string) string {
	var token string
	client := http.Client{}
	var jsonprep string = `{"username":"` + username + `","password":"` + password + `"}`
	var jsonStr = []byte(jsonprep)
	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonStr))
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("body=", string(body))
		data := model.Credenziali{}
		_ = json.Unmarshal([]byte(body), &data)
		token = data.Token

	}
	fmt.Println("token:", token)
	return token
}

func listCustomer(token string) {
	url := "http://api.fintyreclub.it/gommista/users?token=" + token
	fmt.Println("url:", url)
	//res, err := http.Get(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(body)

	var customers = []customer{}
	err = json.Unmarshal([]byte(body), &customers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customers)
}
