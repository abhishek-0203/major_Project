package database

import (
"context"
"fmt"
"log"
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

var (
client *mongo.Client
db     *mongo.Database
)

// InitDB initializes the MongoDB connection
func InitDB() error {
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Connect to MongoDB
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
var err error
client, err = mongo.Connect(ctx, clientOptions)
if err != nil {
return fmt.Errorf("failed to connect to MongoDB: %v", err)
}

// Check the connection
err = client.Ping(ctx, nil)
if err != nil {
return fmt.Errorf("failed to ping MongoDB: %v", err)
}

// Initialize database
db = client.Database("project_management")
log.Println("Connected to MongoDB successfully")
return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
return db.Collection(name)
}

// CloseDB closes the MongoDB connection
func CloseDB() error {
if client == nil {
return nil
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

if err := client.Disconnect(ctx); err != nil {
return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
}

client = nil
db = nil
log.Println("Disconnected from MongoDB")
return nil
}
