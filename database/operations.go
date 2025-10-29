package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UsersCollection      = "users"
	ClientsCollection    = "clients"
	DevelopersCollection = "developers"
	ProjectsCollection   = "projects"
	PaymentsCollection   = "payments"
	ChatCollection       = "chat_messages"
)

// CreateUser creates a new user in the database
func CreateUser(user *User) error {
	user.CreatedAt = time.Now()
	collection := GetCollection(UsersCollection)
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

// FindUserByEmail finds a user by their email address
func FindUserByEmail(email string) (*User, error) {
	var user User
	collection := GetCollection(UsersCollection)
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateClient creates a new client profile
func CreateClient(client *Client) error {
	collection := GetCollection(ClientsCollection)
	_, err := collection.InsertOne(context.Background(), client)
	return err
}

// CreateDeveloper creates a new developer profile
func CreateDeveloper(developer *Developer) error {
	collection := GetCollection(DevelopersCollection)
	_, err := collection.InsertOne(context.Background(), developer)
	return err
}

// CreateProject creates a new project
func CreateProject(project *Project) error {
	project.CreatedAt = time.Now()
	collection := GetCollection(ProjectsCollection)
	_, err := collection.InsertOne(context.Background(), project)
	return err
}

// FindProjects finds all projects matching the filter
func FindProjects(filter bson.M) ([]*Project, error) {
	collection := GetCollection(ProjectsCollection)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var projects []*Project
	if err = cursor.All(context.Background(), &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// CreatePayment creates a new payment record
func CreatePayment(payment *Payment) error {
	payment.CreatedAt = time.Now()
	collection := GetCollection(PaymentsCollection)
	_, err := collection.InsertOne(context.Background(), payment)
	return err
}

// SaveChatMessage saves a new chat message
func SaveChatMessage(message *ChatMessage) error {
	message.CreatedAt = time.Now()
	collection := GetCollection(ChatCollection)
	_, err := collection.InsertOne(context.Background(), message)
	return err
}

// GetChatMessages retrieves chat messages for a project
func GetChatMessages(projectID primitive.ObjectID) ([]*ChatMessage, error) {
	collection := GetCollection(ChatCollection)
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}})
	cursor, err := collection.Find(context.Background(), bson.M{"project_id": projectID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []*ChatMessage
	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
