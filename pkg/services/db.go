package services

import (
    "context"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Ctx = context.TODO()

func InitDB() error {
    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)
    var err error
    Client, err = mongo.Connect(Ctx, clientOptions)
    return err
}


func GetDBClient() *mongo.Client {
    return Client
}
