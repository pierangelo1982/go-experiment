package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/badoux/checkmail"
)

func main() {
	// Open the file
	csvfile, _ := os.Open("mycsv.csv")
	// parse the file
	reader := csv.NewReader(csvfile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if verifyEmail(record[1]) == false {
			fmt.Println(record[1])
		}
	}
}

func verifyEmail(email string) bool {
	//fmt.Println(email)
	err := checkmail.ValidateHost(email)
	if err != nil {
		//fmt.Println(err)
		return false
	} else {
		//fmt.Println("email ok")
		return true
	}
}
