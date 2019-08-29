package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	//"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tealeg/xlsx"
)

var conta int
var numero int

type Anagrafica struct {
	Denominazione string
	Codfisc       string
	Piva          string
	Indirizzo     string
	Cap           string
	Citta         string
	Provincia     string
	Nazione       string
}

func main() {
	readFile()
}

func readFile() {
	excelFileName := "/home/pierangelo/goworkspace/src/github.com/pierangelo1982/go-experiment/importAnagrafiche/clienti.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("errore:", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			conta = conta + 1
			numero = 0
			/*
				marca := row.Cells[0]
				tmpMisura := row.Cells[1]
				misura := stripSpaces(tmpMisura.String()) // tolgo spazi al bianchi tra i codici
				codice := row.Cells[2]
				xl := row.Cells[3]
				nome := row.Cells[4]
				stagione := row.Cells[5]
				// salvo nel db
				//_, err = db.Exec("INSERT INTO pneumatici (marca, misura, codice, xl, nome, stagione) VALUES" + fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s')", marca, misura, codice, xl, nome, stagione))
				if err != nil {
					panic(err)
				}


				fmt.Printf("%s \t %s \t %s \t %s \t %s \t %s \n", marca, misura, codice, xl, nome, stagione)
			*/
			fmt.Println(row.Cells[0])
			fmt.Println("CONTATORE:", conta)
			fmt.Println("---------------------------------------------------------------------------------")

			// Set client options
			clientOptions := options.Client().ApplyURI("mongodb://92.222.77.106:27018")
			//clientOptions := mongo.NewClient(options.Client().ApplyURI("mongodb://92.222.77.106:27018"))
			// Connect to MongoDB
			client, err := mongo.Connect(context.TODO(), clientOptions)

			if err != nil {
				log.Fatal(err)
			}

			// Check the connection
			err = client.Ping(context.TODO(), nil)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Connected to MongoDB!")

			collection := client.Database("anagrafiche").Collection("registries")

			ruan := Anagrafica{row.Cells[0].String(), row.Cells[1].String(), row.Cells[2].String(), row.Cells[3].String(), row.Cells[4].String(), row.Cells[5].String(), row.Cells[6].String(), row.Cells[7].String()}

			insertResult, err := collection.InsertOne(context.TODO(), ruan)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

		}
	}
}

// eliminate white spaces from strings
func stripSpaces(words string) string {
	x := strings.Replace(words, " ", "", -1)
	return x
}
