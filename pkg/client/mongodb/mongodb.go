package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (db *mongo.Database, err error) {
	var mongoDbURL string
	var isAuth bool
	if username == "" && password == "" {
		mongoDbURL = fmt.Sprintf("mongodb//%s:%s", host, port)
	} else {
		mongoDbURL = fmt.Sprintf("mongodb//%s:%s@%s:%s", username, password, host, port)
		isAuth = true
	}

	options.Client().ApplyURI(mongoDbURL).SetAuth(options.Credential{})

	clientOptions := options.Client().ApplyURI(mongoDbURL)
	if isAuth {
		if authDb == "" {
			authDb = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDb,
			Username:   username,
			Password:   password,
		})

	}

	// Connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongoDB due to error: %v", err)
	}

	//Ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client.Database(database), nil
}
