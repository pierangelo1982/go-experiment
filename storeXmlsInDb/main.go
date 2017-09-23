package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"

	_ "github.com/go-sql-driver/mysql"
)

var conta int
var numero int

var marca string
var misura string
var codice string
var xl string
var nome string
var stagione string

type dbConn struct {
	db *sql.DB
}

func main() {
	connectDB()
	readFile()
}

func readFile() {
	// db connection
	db, err := sql.Open("mysql", "root:alnitek@tcp(0.0.0.0:3308)/gociao")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	excelFileName := "/home/pierangelo/goworkspace/src/github.com/pierangelo1982/go-experiment/storeXmlsInDb/00-UNITI.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("errore:", err)
	}
	for _, sheet := range xlFile.Sheets[:1] {
		for _, row := range sheet.Rows {
			conta = conta + 1
			numero = 0

			marca := row.Cells[0]
			tmpMisura := row.Cells[1]
			misura := stripSpaces(tmpMisura.String()) // tolgo spazi al bianchi tra i codici
			codice := row.Cells[2]
			xl := row.Cells[3]
			nome := row.Cells[4]
			stagione := row.Cells[5]
			// salvo nel db
			_, err = db.Exec("INSERT INTO pneumatici (marca, misura, codice, xl, nome, stagione) VALUES" + fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s')", marca, misura, codice, xl, nome, stagione))
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s \t %s \t %s \t %s \t %s \t %s \n", marca, misura, codice, xl, nome, stagione)
			//fmt.Println("CONTATORE:", conta)
			fmt.Println("---------------------------------------------------------------------------------")
		}
	}
	println("TOTALE:", conta)
	println("range sheet:", xlFile.Sheets)
}

func connectDB() {
	db, err := sql.Open("mysql", "root:alnitek@tcp(0.0.0.0:3308)/gociao")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	// create db se non esiste
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + "gociao")
	if err != nil {
		panic(err)
	}
	// creare tabella
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS pneumatici ( id integer NOT NULL AUTO_INCREMENT,marca varchar(250) null, misura varchar(250) null, codice varchar(250) null, xl varchar(250) null, nome varchar(250) null, stagione varchar(250) null, PRIMARY KEY (id) )")
	if err != nil {
		panic(err)
	}
}

// eliminate white spaces from strings
func stripSpaces(words string) string {
	x := strings.Replace(words, " ", "", -1)
	return x
}
