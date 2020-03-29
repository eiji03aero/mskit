package eventstore

import (
	"context"
	"encoding/json"

	"github.com/eiji03aero/mskit"
	mongodb "github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Collection_Events = "events"
)

type DBOption struct {
	Host string
	Port string
	Name string
}

type Client struct {
	Database      string
	Client        *mongo.Client
	eventRegistry *mskit.EventRegistry
}

func New(opt mongodb.DBOption, er *mskit.EventRegistry) (mskit.EventStore, error) {
	if opt.Name == "" {
		opt.Name = mongodb.DefaultDatabaseName
	}

	client, err := mongodb.GetClient(opt)
	if err != nil {
		return nil, err
	}

	es := &Client{
		Database:      opt.Name,
		Client:        client,
		eventRegistry: er,
	}

	return es, nil
}

func (c *Client) Save(event mskit.Event) error {
	col := c.collection()

	eventDoc := NewEventDocument(event)
	err := eventDoc.makeJsonData()
	if err != nil {
		return err
	}

	_, err = col.InsertOne(context.Background(), eventDoc)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Load(id string, aggregate mskit.Aggregate) error {
	col := c.collection()
	aggregateName := utils.GetTypeName(aggregate)

	cur, err := col.Find(
		context.Background(),
		bson.D{
			{Key: "aggregate_type", Value: aggregateName},
			{Key: "aggregate_id", Value: id},
		},
	)
	if err != nil {
		return err
	}

	for cur.Next(context.Background()) {
		var eventDoc EventDocument
		err := cur.Decode(&eventDoc)
		if err != nil {
			return err
		}

		eventPtr, err := c.eventRegistry.Get(eventDoc.Type)
		if err != nil {
			return err
		}

		err = json.Unmarshal(eventDoc.JsonData, eventPtr)
		if err != nil {
			return err
		}

		event := utils.DereferenceIfPtr(eventPtr)
		err = aggregate.Apply(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) db() *mongo.Database {
	return c.Client.Database(c.Database)
}

func (c *Client) collection() *mongo.Collection {
	return c.db().Collection(Collection_Events)
}
