package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", tyres)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func tyres(w http.ResponseWriter, r *http.Request) {
	// db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:<port>)/<dbname>?charset=utf8" )
	db, err := sql.Open("mysql", "root:alnitek@tcp(0.0.0.0:3308)/gociao?charset=utf8")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from pneumatici")

	if err != nil {
		log.Fatal(err)
	}
	type Tyre struct {
		ID       int    `json:"id"`
		Marca    string `json:"marca"`
		Misura   string `json:"misura"`
		Codice   string `json:"codice"`
		Xl       string `json:"xl"`
		Nome     string `json:"nome"`
		Stagione string `json:"stagione"`
	}

	var tyres []Tyre

	for rows.Next() {
		var id int
		var marca string
		var misura string
		var codice string
		var xl string
		var nome string
		var stagione string

		rows.Scan(&id, &marca, &misura, &codice, &xl, &nome, &stagione)
		tyres = append(tyres, Tyre{id, marca, misura, codice, xl, nome, stagione})
	}

	tyresBytes, _ := json.Marshal(&tyres)

	w.Write(tyresBytes)
	db.Close()
}
