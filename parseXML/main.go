package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Word struct {
	Di   string `xml:"di,attr"`
	Ga   string `xml:"ga,attr"`
	Anno string `xml:"yravg,attr"`
}

type Chapter struct {
	Number string `xml:"n,attr"`
	Words  []Word `xml:"w"`
}
type Book struct {
	Number string `xml:"n,attr"`
	//Chapter  string    `xml:"c"`
	Chapters []Chapter `xml:"c"`
}

type Query struct {
	Books []Book `xml:"b"`
}

func main() {
	// Open our xmlFile
	xmlFile, err := os.Open("index.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var query Query
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &query)

	//fmt.Println(query.Books)

	for i := 0; i < len(query.Books); i++ {
		fmt.Println("Libri:", query.Books[i].Number)
		for x := 0; x < len(query.Books[i].Chapters); x++ {
			fmt.Println("Capitoli:", query.Books[i].Chapters[x].Number)
			for y := 0; y < len(query.Books[i].Chapters[x].Words); y++ {
				fmt.Println("id", query.Books[i].Chapters[x].Words[y].Di)
				fmt.Println("ga", query.Books[i].Chapters[x].Words[y].Ga)
				fmt.Println("anno", query.Books[i].Chapters[x].Words[y].Anno)
				fmt.Println("+++")
			}

			fmt.Println("----------------------------------------------")
		}
	}

}
