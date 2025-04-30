package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionUser = "users"

type RoleType string

const (
	AdminRole RoleType = "admin"
	UserRole  RoleType = "user"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      RoleType           `json:"role" bson:"role"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserRepository interface {
	RegisterUser(ctx context.Context, user *User) (string, error)
	LoginUser(ctx context.Context, loginRequest LoginRequest) (string, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, user *User) (string, error)
	LoginUser(ctx context.Context, loginRequest LoginRequest) (string, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}
