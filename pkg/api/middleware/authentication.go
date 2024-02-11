package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to authenticate requests using bearer token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        // Check if the authorization header starts with "Bearer "
        if !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }

        // Extract the token from the authorization header
        token := strings.TrimPrefix(authHeader, "Bearer ")

        // Validate the token (e.g., verify it against a database or external service)
        valid := isValidToken(token)
        if !valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// isValidToken is a placeholder function to validate the bearer token (dummy token)
func isValidToken(token string) bool {
    return true
}
