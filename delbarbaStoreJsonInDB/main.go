package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pierangelo1982/go-experiment/delbarbaStoreJsonInDB/model"
	"github.com/pierangelo1982/go-experiment/delbarbaStoreJsonInDB/utils"
)

const username string = "xxxx"
const password string = "xxxx"
const loginURL = "http://api.fintyreclub.it/gommista/token"

var myToken string

func main() {
	fmt.Println(username)
	fmt.Println(password)
	myToken := getToken(loginURL, username, password)
	listCustomer(myToken)
	//fmt.Println(myToken)
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
		//fmt.Println("body=", string(body))
		data := model.Credenziali{}
		_ = json.Unmarshal([]byte(body), &data)
		token = data.Token

	}
	fmt.Println("token:", token)
	return token
}

func listCustomer(token string) {
	url := "http://api.fintyreclub.it/gommista/users?token=" + token
	//fmt.Println("url:", url)
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
	//fmt.Println(body)

	var customers = []model.Customer{}
	err = json.Unmarshal([]byte(body), &customers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customers)

	db, err := sql.Open("mysql", "root:alnitek82@tcp(0.0.0.0:3310)/delbarba_backend_development")
	if err != nil {
		panic(err.Error())
	}

	//myDateTime := time.Now().Format("2006-01-02 15:04:05")
	for i := range customers {
		isAPP := 1
		_, err = db.Exec("INSERT INTO customers (customer_code, name, surname, id_gommista, address, city, prov, zip, phone, email, mobile, note, driving_license_expiration, is_app) VALUES" + fmt.Sprintf("(%s, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d')", customers[i].ID, customers[i].Nome, customers[i].Cognome, customers[i].IDGommista, customers[i].Indirizzo, customers[i].Citta, customers[i].Provincia, customers[i].Cap, customers[i].Telefono, customers[i].Email, customers[i].Cellulare, customers[i].Note, utils.ParseData(customers[i].DataScadenzaPatente).Format("2006-01-02"), isAPP))
		if err != nil {
			fmt.Println(err)
		}
	}
	db.Close()
	//fmt.Println(myDateTime)
}
