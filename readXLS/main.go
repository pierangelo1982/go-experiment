package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var conta int
var numero int

func main() {
	readFile()
}

func readFile() {
	excelFileName := "/home/pierangelo/goworkspace/src/github.com/pierangelo1982/go-experiment/readXLS/00-UNITI.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("errore:", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			conta = conta + 1
			numero = 0
			for _, cell := range row.Cells {
				text := cell.String()
				numero = numero + 1
				fmt.Print(numero)
				fmt.Printf("\t %s\n", text)
			}
			fmt.Println("CONTATORE:", conta)
			fmt.Println("----------------------------------------------------------")
		}
	}
	println("range sheet:", xlFile.Sheets)
}
