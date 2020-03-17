package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultDatabaseName = "mskit"
)

type DBOption struct {
	Host string
	Port string
	Name string
}

func GetDBUrl(opt DBOption) string {
	return fmt.Sprintf(
		"mongodb://%s:%s",
		opt.Host,
		opt.Port,
	)
}

func GetClient(opt DBOption) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return mongo.Connect(ctx, options.Client().ApplyURI(GetDBUrl(opt)))
}
