package user

import (
	"context"
	"fmt"
	"majorProject/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UsersCollection = "users"
)

// User represents the user model in MongoDB
type User struct {
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"password_hash"`
	Role         string    `bson:"role"`
	Name         string    `bson:"name"`
	CreatedAt    time.Time `bson:"created_at"`
}

// CreateUser creates a new user in MongoDB
func CreateUser(user *User) error {
	collection := database.GetCollection(UsersCollection)
	user.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if email already exists
	count, err := collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return fmt.Errorf("error checking email existence: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists")
	}

	// Insert the user
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

// FindUserByEmail finds a user by their email address
func FindUserByEmail(email string) (*User, error) {
	collection := database.GetCollection(UsersCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error finding user: %v", err)
	}

	return &user, nil
}
