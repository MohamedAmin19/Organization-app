// services/organizations.go
package services

import (
    "errors"
    "structure/pkg/database/mongodb"
    "structure/pkg/database/mongodb/models"
)

// CreateOrganization creates a new organization in the database
func CreateOrganization(name, description string) (string, error) {
    // Create a new organization object
    organization := models.Organization{
        Name:        name,
        Description: description,
    }

    // Insert the organization into the database
    insertedID, err := mongodb.InsertOneOrganization("organizations", organization)
    if err != nil {
        return "", err
    }

    return insertedID, nil
}

// GetOrganizationByIDWithMembers retrieves an organization by its ID along with its members from the database
func GetOrganizationByIDWithMembers(organizationID string) (*models.Organization, error) {
    organization, err := mongodb.GetOrganizationByIDWithMembers(organizationID)
    if err != nil {
        return nil, err
    }
    return organization, nil
}

// GetAllOrganizations retrieves all organizations from the database
func GetAllOrganizations() ([]models.Organization, error) {
    // Query the database to get all organizations
    organizations, err := mongodb.GetAllOrganizations()
    if err != nil {
        return nil, err
    }

    return organizations, nil
}

func UpdateOrganization(organizationID, name, description string) (*models.Organization, error) {
    // Call the database operation to update the organization
    updatedOrganization, err := mongodb.UpdateOrganization(organizationID, name, description)
    if err != nil {
        return nil, err
    }

    return updatedOrganization, nil
}

// DeleteOrganizationByID deletes an organization by its ID from the database
func DeleteOrganizationByID(organizationID string) error {
    // Call the corresponding MongoDB function to delete the organization
    err := mongodb.DeleteOrganizationByID(organizationID)
    return err
}

// InviteUserToOrganization sends an invitation to a user to join an organization
func InviteUserToOrganization(organizationID, userEmail string) error {
    // Retrieve organization details
    organization, err := mongodb.GetOrganizationByIDWithMembers(organizationID)
    if err != nil {
        return err
    }
    if organization == nil {
        return errors.New("Organization not found")
    }

    // Save the invitation in the database
    invitation := models.Invitation{
        OrganizationID: organizationID,
        UserEmail:      userEmail,
        Sent:           true,
    }
    err = mongodb.SaveInvitation(invitation)
    if err != nil {
        return err
    }

    return nil
}