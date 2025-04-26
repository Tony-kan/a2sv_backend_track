package services

import (
	"context"
	"fmt"
	"log"
	"task_6_task_management_api_with_auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	RegisterUser(user models.User) (string, error)
	LoginUser(loginRequest models.LoginRequest) (string, error)
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	userCollection *mongo.Collection
}

func NewUserService(userCollection *mongo.Collection) UserServices {
	ctx := context.Background()
	// indexModel := mongo.IndexModel{
	// 	Keys: bson.D{{"createdAt", 1}},
	// }
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("uniq_username"),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("uniq_email"),
		},
	}

	if _, err := userCollection.Indexes().CreateMany(ctx, indexModels); err != nil {
		panic(fmt.Sprintf("Failed to create indexes: %v", err))
	}
	return &UserService{
		userCollection: userCollection,
	}
}

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("secret")

func (service *UserService) RegisterUser(user models.User) (string, error) {

	if user.Email == "" || user.Password == "" || user.Username == "" {
		return "", fmt.Errorf("username,email and password are required")
	}

	user.ID = primitive.NewObjectID()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if user.Role == "" {
		user.Role = models.UserRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	result, err := service.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		if isDupKey(err) {
			return "", fmt.Errorf("username or email already exists")
		}
		return "", fmt.Errorf("failed to create user: %v", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (service *UserService) LoginUser(loginRequest models.LoginRequest) (string, error) {
	log.Println("Jwt secret ------------------------", jwtSecret)
	var user models.User
	err := service.userCollection.FindOne(context.Background(), bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		return "", fmt.Errorf("failed to find user with email: %v ", loginRequest.Email)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return "", fmt.Errorf("invalid credentials: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", fmt.Errorf("failed to create token: %v", err)
	}

	return jwtToken, nil
}

func (service *UserService) GetAllUsers() ([]models.User, error) {

	users := make([]models.User, 0)

	cur, err := service.userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		// var task models.Task
		var user models.User

		if err := cur.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode task: %v", err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
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
