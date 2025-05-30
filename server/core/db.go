package core

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBConnection struct {
	client   *mongo.Client
	db       *mongo.Database
	todoColl *mongo.Collection
	userColl *mongo.Collection
}

func InitDB(connectionString string) (*DBConnection, error) {
	clientSettings := options.Client().ApplyURI(connectionString)
	clientSettings.SetAppName("todo-tui-server")

	client, err := mongo.Connect(clientSettings)

	if err != nil {
		return nil, err
	}

	db := client.Database("todo")

	conn := DBConnection{
		client:   client,
		db:       db,
		todoColl: db.Collection("todo"),
		userColl: db.Collection("users"),
	}

	return &conn, nil
}

func (self *DBConnection) Disconnect(ctx context.Context) error {
	return self.client.Disconnect(ctx)
}
