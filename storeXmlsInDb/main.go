package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var conta int
var numero int

var marca string
var misura string
var codice string
var xl string
var nome string
var stagione string

func main() {
	readFile()
}

func readFile() {
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
			misura := row.Cells[1]
			codice := row.Cells[2]
			xl := row.Cells[3]
			nome := row.Cells[4]
			stagione := row.Cells[5]
			fmt.Printf("%s \t %s \t %s \t %s \t %s \t %s \n", marca, misura, codice, xl, nome, stagione)
			//fmt.Println("CONTATORE:", conta)
			fmt.Println("---------------------------------------------------------------------------------")
		}
	}
	println("TOTALE:", conta)
	println("range sheet:", xlFile.Sheets)
}
