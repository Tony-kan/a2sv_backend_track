package repositories

import (
	"context"
	"fmt"
	"strings"
	"task_7_task_management_api_refactoring/domain"
	"task_7_task_management_api_refactoring/infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(database *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}

func (repository *userRepository) GetCollection() *mongo.Collection {
	return repository.database.Collection(repository.collection)
}
func (repository *userRepository) RegisterUser(ctx context.Context, user *domain.User) (string, error) {

	if strings.TrimSpace(user.Password) == "" {
		return "", domain.ErrEmptyPassword
	}

	if user.Email == "" || user.Username == "" {
		return "", fmt.Errorf("username and email are required")
	}

	hashedPassword, err1 := infrastructure.HashPassword(user.Password)
	if err1 != nil {
		return "", fmt.Errorf("failed to hash password: %v", err1)
	}

	user.ID = primitive.NewObjectID()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	if user.Role == "" {
		user.Role = "user"
	}

	user.Password = hashedPassword

	_, err := repository.GetCollection().InsertOne(ctx, user)
	if err != nil {
		if isDupKey(err) {
			return "", domain.ErrUserExists
		}
		return "", fmt.Errorf("failed to create user: %v", err)
	}
	return user.ID.Hex(), nil
}
func (repository *userRepository) LoginUser(ctx context.Context, loginRequest domain.LoginRequest) (string, error) {
	var user domain.User
	err := repository.GetCollection().FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", domain.ErrUserNotFound
		}
		return "", fmt.Errorf("failed to login user: %v", err)
	}

	// Check if the password is correct
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
	// 	return "", fmt.Errorf("invalid credentials: %v", err)
	// }
	if err := infrastructure.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return "", fmt.Errorf("invalid credentials: %v", err)
	}

	// Generate JWT token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user_id": user.ID,
	// 	"email":   user.Email,
	// 	"role":    user.Role,
	// 	"exp":     time.Now().Add(72 * time.Hour).Unix(),
	// })

	// jwtToken, err := token.SignedString(jwtSecret)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to create token: %v", err)
	// }
	// return jwtToken, nil
	token, err := infrastructure.GenerateJWTToken(user.ID.Hex(), user.Email, string(user.Role))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %v", err)
	}
	return token, nil
}

func (repository *userRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := repository.GetCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %v", err)
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	return users, nil
}
func isDupKey(err error) bool {
	if we, ok := err.(mongo.WriteException); ok {
		for _, e := range we.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}
