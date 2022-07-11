package config

import (
	"context"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *ConfigImpl) MongoDB(uri, database string) (*mongo.Client, *mongo.Database) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().SetMonitor(apmmongo.CommandMonitor()).ApplyURI(uri),
	)
	if err != nil {
		c.Logger.Panic(err)
	}
	return client, client.Database(database)
}
