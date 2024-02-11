package models

// SignInData represents the structure of data for sign-in requests
type SignInData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}