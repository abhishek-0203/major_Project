package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user model in MongoDB
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Role         string             `bson:"role"`
	CreatedAt    time.Time          `bson:"created_at"`
}

// Client represents the client profile
type Client struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Name        string             `bson:"name"`
	CompanyName string             `bson:"company_name,omitempty"`
	Phone       string             `bson:"phone"`
}

// Developer represents the developer profile
type Developer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Name       string             `bson:"name"`
	Skills     []string           `bson:"skills"`
	Experience int                `bson:"experience"`
	Portfolio  string             `bson:"portfolio,omitempty"`
}

// Project represents a project in the system
type Project struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	ClientID    primitive.ObjectID   `bson:"client_id"`
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	Budget      float64              `bson:"budget"`
	Status      string               `bson:"status"`
	CreatedAt   time.Time            `bson:"created_at"`
	Developers  []primitive.ObjectID `bson:"developers,omitempty"`
}

// Payment represents a payment transaction
type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `bson:"project_id"`
	Amount    float64            `bson:"amount"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
}

// ChatMessage represents a message in the chat system
type ChatMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `bson:"project_id"`
	SenderID  primitive.ObjectID `bson:"sender_id"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}
