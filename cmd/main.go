package main

import (
    "github.com/joho/godotenv"
    "structure/pkg/api/routes"
    "structure/pkg/services"
    "log"
)

func main() {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize the database connection
    if err := services.InitDB(); err != nil {
        panic(err)
    }
    defer services.Client.Disconnect(services.Ctx)

    // Setup the router and start the server
    router := routes.SetupRouter()
    router.Run(":8080")
}
