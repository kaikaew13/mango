package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	MONGO_URI := os.Args[1]

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Disconnect(ctx)
	fmt.Println("connected successfully")

	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, db := range dbs {
		fmt.Println(db)
	}
}
