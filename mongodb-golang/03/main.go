package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pierangelo1982/go-experiment/mongodb-golang/03/config"
)

type agenda struct {
	Nome    string
	Cognome string
	Nazione string
}

func main() {
	cur, err := config.ConnectCollection()
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		if err = cur.Decode(elem); err != nil {
			fmt.Errorf("readTasks: couldn't make to-do item ready for display: %v", err)
		}
		fmt.Println(elem.Lookup("nome").StringValue())
	}
}
