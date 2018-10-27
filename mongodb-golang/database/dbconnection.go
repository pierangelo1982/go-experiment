package database

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func DbConnection() {
	type anagrafica struct {
		Nome    string
		Cognome string
	}
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("demo").Collection("agenda")
	fmt.Println(collection)
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem.)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
