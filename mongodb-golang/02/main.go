package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type agenda struct {
	Nome    string
	Cognome string
	Nazione string
}

func main() {
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
		log.Fatalln(err)
	}
	defer cur.Close(context.Background())
	fmt.Println("ciao", cur)

	tw := tabwriter.NewWriter(os.Stdout, 24, 2, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintln(tw, "Created At\tModified At\tTask\t")

	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		if err = cur.Decode(elem); err != nil {
			fmt.Errorf("readTasks: couldn't make to-do item ready for display: %v", err)
		}
		//fmt.Println(elem.Lookup("nome").StringValue())
		t := agenda{
			Nome:    elem.Lookup("nome").StringValue(),
			Cognome: elem.Lookup("cognome").StringValue(),
			Nazione: elem.Lookup("nazione").StringValue(),
		}
		output := fmt.Sprintf("%s\t%s\t%s\t",
			t.Nome,
			t.Cognome,
			t.Nazione,
		)
		fmt.Fprintln(tw, output)
		if err = tw.Flush(); err != nil {
			fmt.Errorf("readTasks: all data for the to-do couldn't be printed: %v", err)
		}

	}
}
