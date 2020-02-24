package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventstore/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	dbOptions := mongo.DBOption{
		Host: "ftgo-kitchen-mongo",
		Port: "27017",
	}

	er := mskit.NewEventRegistry()

	eventStore, err := mongo.New(dbOptions, er)
	if err != nil {
		panic(err)
	}

	// _ = mskit.NewRepository(eventStore)

	mongoClient, ok := eventStore.(*mongo.Client)
	if !ok {
		panic("shippai mongo client")
	}

	db := mongoClient.Client.Database("mskit")
	coll := db.Collection("testing")
	coll.Drop(context.Background())

	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{Key: "item", Value: "canvas"},
			{Key: "quantity", Value: 100},
			{Key: "tags", Value: bson.A{"cotton"}},
		},
	)

	id := result.InsertedID
	log.Println(id, err)

	var doc bson.D
	coll.FindOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	).Decode(&doc)
	docStr, err := json.Marshal(doc)
	log.Println(string(docStr), err)
}
