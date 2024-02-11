package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Organization represents an organization entity in the database
type Organization struct {
    ID                  primitive.ObjectID `bson:"_id,omitempty" json:"organization_id"`
    Name                string             `bson:"name" json:"name"`
    Description         string             `bson:"description" json:"description"`
    OrganizationMembers []OrganizationMember `bson:"organization_members" json:"organization_members"`
}

type OrganizationMember struct {
    Name        string `bson:"name" json:"name"`
    Email       string `bson:"email" json:"email"`
    AccessLevel string `bson:"access_level" json:"access_level"`
}