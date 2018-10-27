package config

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func ConnectionDB() *mongo.Client {
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func ConnectCollection() (mongo.Cursor, error) {
	collection := ConnectionDB().Database("demo").Collection("agenda")

	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	return cur, nil
}
