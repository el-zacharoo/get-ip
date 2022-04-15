package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	locaColl *mongo.Collection
}

func Connect() *Store {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	mongoAPI := os.Getenv("API")
	fmt.Println(mongoAPI)

	clientOptions := options.Client().ApplyURI("mongodb+srv://geoloaction:e2Fyk5w2ZJnV6uzN@cluster0.u4qeu.mongodb.net")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("info")

	return &Store{
		locaColl: db.Collection("location"),
	}
}
