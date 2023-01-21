package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", viper.Get("mongodb.user"), viper.Get("mongodb.password"), viper.Get("mongodb.host"), viper.Get("mongodb.port"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(fmt.Sprintf("mongodb connect error [%s], err = %s", uri, err))
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("mongodb ping error [%s], err = %s", uri, err))
		return nil, ctx, err
	}
	fmt.Printf("Successful connection to mongodb: [%s]", uri)
	return client, ctx, err
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("communication_service").Collection(collectionName)
	return collection
}
