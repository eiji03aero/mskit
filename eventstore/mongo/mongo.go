package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/eiji03aero/mskit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBOption struct {
	Host string
	Port string
}

type Client struct {
	Client        *mongo.Client
	eventRegistry *mskit.EventRegistry
}

func getDBUrl(opt DBOption) string {
	return fmt.Sprintf(
		"mongodb://%s:%s",
		opt.Host,
		opt.Port,
	)
}

func New(opt DBOption, er *mskit.EventRegistry) (mskit.EventStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getDBUrl(opt)))
	if err != nil {
		return nil, err
	}

	es := &Client{
		Client:        client,
		eventRegistry: er,
	}

	return es, nil
}

func (c *Client) Save(event *mskit.Event) error {
	return nil
}

func (c *Client) Load(id string, aggregate mskit.Aggregate) error {
	return nil
}
