package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "structure/pkg/services"
    "structure/pkg/database/mongodb"
)

// Creates a new organization
func CreateOrganization(c *gin.Context) {
    // Parse the request body
    var requestBody struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create the organization
    organizationID, err := services.CreateOrganization(requestBody.Name, requestBody.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
        return
    }

    // Return the organization ID in the response
    c.JSON(http.StatusOK, gin.H{"organization_id": organizationID})
}

// Get an organization by Id
func GetOrganizationByID(c *gin.Context) {
    organizationID := c.Param("organization_id")

    // Query the database for organization details along with its members
    organization, err := mongodb.GetOrganizationByIDWithMembers(organizationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organization"})
        return
    }
    if organization == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
        return
    }

    // Return the fetched organization details in the response
    c.JSON(http.StatusOK, organization)
}

// Retrieves all organizations
func GetAllOrganizations(c *gin.Context) {
    // Get all organizations from the database
    organizations, err := services.GetAllOrganizations()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organizations"})
        return
    }

    // Return the organizations in the response
    c.JSON(http.StatusOK, organizations)
}

// Update an organization by Id
func UpdateOrganization(c *gin.Context) {
    organizationID := c.Param("organization_id")

    // Parse the request body
    var requestBody struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the organization
    updatedOrganization, err := services.UpdateOrganization(organizationID, requestBody.Name, requestBody.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organization"})
        return
    }

    // Return the updated organization details in the response
    c.JSON(http.StatusOK, updatedOrganization)
}

// Delete an organization by Id
func DeleteOrganization(c *gin.Context) {
    organizationID := c.Param("organization_id")

    // Delete the organization
    err := services.DeleteOrganizationByID(organizationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organization"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

// Invite to an organization by Id
func InviteUserToOrganization(c *gin.Context) {
    organizationID := c.Param("organization_id")

    var requestBody struct {
        UserEmail string `json:"user_email"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Call service to send invitation
    err := services.InviteUserToOrganization(organizationID, requestBody.UserEmail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send invitation"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Invitation sent successfully"})
}