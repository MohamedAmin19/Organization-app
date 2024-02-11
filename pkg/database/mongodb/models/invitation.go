package models

// Invitation represents an invitation to join an organization
type Invitation struct {
    OrganizationID string `bson:"organization_id" json:"organization_id"`
    UserEmail      string `bson:"user_email" json:"user_email"`
    Sent           bool   `bson:"sent" json:"sent"`
}
