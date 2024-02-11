package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "structure/pkg/database/mongodb/models"
    "structure/pkg/services"
)

//SignUp a new user
func SignUp(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the user's password before storing it in the database
    hashedPassword, err := services.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = hashedPassword

    collection := services.GetDBClient().Database("organization-app").Collection("users")

    _, err = collection.InsertOne(services.Ctx, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

//SignIn an existing user
func SignIn(c *gin.Context) {
    var signInData models.SignInData
    if err := c.BindJSON(&signInData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Fetch the user from the database based on the provided email
    user, err := services.GetUserByEmail(signInData.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Check if the provided password matches the hashed password stored in the database
    if !services.CheckPasswordHash(signInData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT tokens (access token and refresh token)
    accessToken, refreshToken, err := services.GenerateTokens(user.ID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User signed in successfully",
        "access_token": accessToken,
        "refresh_token": refreshToken,
    })
}

//Refresh token
func RefreshToken(c *gin.Context) {
    var refreshTokenData struct {
        RefreshToken string `json:"refresh_token"`
    }
    if err := c.BindJSON(&refreshTokenData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Generate new access and refresh tokens based on the provided refresh token
    accessToken, refreshToken, err := services.RefreshTokens(refreshTokenData.RefreshToken)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh tokens"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":       "Token refreshed successfully",
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
}


