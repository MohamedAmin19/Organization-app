// services/auth.go
package services

import (
    "time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"structure/pkg/database/mongodb/models"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("B82EExjz8WdS8oXomc6kX7V1BfijZWTvXXoLpwgJXJw=") //example
const bcryptCost = 10 //cost factor


// GenerateTokens generates access and refresh tokens for a given user ID
func GenerateTokens(userID string) (accessToken string, refreshToken string, err error) {
    // Create access token
    accessToken, err = createToken(userID, time.Hour*1)
    if err != nil {
        return "", "", err
    }

    // Create refresh token
    refreshToken, err = createToken(userID, time.Hour*24*7)
    if err != nil {
        return "", "", err
    }

    return accessToken, refreshToken, nil
}

// RefreshTokens refreshes the access token using the refresh token
func RefreshTokens(refreshToken string) (accessToken string, newRefreshToken string, err error) {
    // Parse the refresh token
    userID, err := parseToken(refreshToken)
    if err != nil {
        return "", "", err
    }

    // Create new access token
    accessToken, err = createToken(userID, time.Hour*1)
    if err != nil {
        return "", "", err
    }

    // Create new refresh token (optional)
    newRefreshToken, err = createToken(userID, time.Hour*24*7)
    if err != nil {
        return "", "", err
    }

    return accessToken, newRefreshToken, nil
}

//create a JWT token
func createToken(userID string, expirationTime time.Duration) (string, error) {
    // Create JWT claims
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(expirationTime).Unix(),
    }

    // Create token object with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with secret key
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// parse and validate JWT token
func parseToken(tokenString string) (userID string, err error) {
    // Parse and validate token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return "", err
    }

    // Extract userID from token claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", err
    }

    userID, ok = claims["userID"].(string)
    if !ok {
        return "", err
    }

    return userID, nil
}

func GetUserByEmail(email string) (*models.User, error) {
    if MockGetUserByEmail != nil {
        return MockGetUserByEmail(email)
    }

    // Get a MongoDB collection
    collection := Client.Database("organization-app").Collection("users")

    // Define a filter to find the user by email
    filter := bson.M{"email": email}

    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Perform the query to find the user
    var user models.User
    err := collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

// HashPassword generates a bcrypt hash of the provided password
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash checks if a given password matches the hashed password stored in the database for a specific user
func CheckPasswordHash(password, hashedPassword string) bool {
    // Compare the provided password with the hashed password using bcrypt
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

