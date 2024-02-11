// services/service_mock.go
package services

import (
    "structure/pkg/database/mongodb/models"
)

// MockGetUserByEmail mocks the GetUserByEmail function for testing purposes
var MockGetUserByEmail func(email string) (*models.User, error)
