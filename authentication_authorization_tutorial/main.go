package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]*User)

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("your_jwt_secret")

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to the Authentication and Authorization Tutorial "})
	})

	// Register
	router.POST("/register", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		//Todo : Implement user registration logic
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		user.Password = string(hashedPassword)
		users[user.Email] = &user

		ctx.JSON(200, gin.H{"message": "User registered successfully"})
	})

	// Login
	router.POST("/login", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		// Todo : Implement user login logic
		existingUser, ok := users[user.Email]

		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
			ctx.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": existingUser.ID,
			"email":   existingUser.Email,
		})

		jwtToken, err := token.SignedString(jwtSecret)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "User logged in Successful", "jwtToken": jwtToken})
	})

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Users list", "users": users})
	})

	router.Run(":8080")
}
