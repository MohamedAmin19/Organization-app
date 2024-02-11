package mongodb

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"structure/pkg/database/mongodb/models"
)

const DatabaseName = "organization-app"

// GetMongoDBClient returns a MongoDB client instance
func GetMongoDBClient() (*mongo.Client, error) {
    // Set MongoDB connection options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    // Ping the MongoDB server to verify that the connection was successful
    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    return client, nil
}

// InsertOneOrganization inserts a single organization document into the specified collection
func InsertOneOrganization(collectionName string, document interface{}) (string, error) {
    // Connect to the MongoDB client
    client, err := GetMongoDBClient()
    if err != nil {
        return "", err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection(collectionName)

    // Insert the document into the collection
    result, err := collection.InsertOne(context.Background(), document)
    if err != nil {
        return "", err
    }

    // Get the ID of the inserted document
    insertedID := result.InsertedID.(primitive.ObjectID).Hex()

    return insertedID, nil
}

// GetOrganizationByIDWithMembers retrieves an organization by its ID along with its members from the database
func GetOrganizationByIDWithMembers(organizationID string) (*models.Organization, error) {
    // Connect to MongoDB
    client, err := GetMongoDBClient()
    if err != nil {
        return nil, err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection("organizations")

    // Convert the organization ID string to MongoDB ObjectID
    objID, err := primitive.ObjectIDFromHex(organizationID)
    if err != nil {
        return nil, err
    }

    // Find the organization by its ID
    var organization models.Organization
    err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&organization)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }

    return &organization, nil
}

// GetAllOrganizations retrieves all organizations from the database
func GetAllOrganizations() ([]models.Organization, error) {
    // Connect to MongoDB
    client, err := GetMongoDBClient()
    if err != nil {
        return nil, err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection("organizations")

    // Find all organizations
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var organizations []models.Organization
    if err := cursor.All(context.Background(), &organizations); err != nil {
        return nil, err
    }

    return organizations, nil
}

func UpdateOrganization(organizationID, name, description string) (*models.Organization, error) {
    // Connect to MongoDB
    client, err := GetMongoDBClient()
    if err != nil {
        return nil, err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection("organizations")

    // Convert the organization ID string to MongoDB ObjectID
    objID, err := primitive.ObjectIDFromHex(organizationID)
    if err != nil {
        return nil, err
    }

    // Prepare update data
    updateData := bson.M{
        "$set": bson.M{
            "name":        name,
            "description": description,
        },
    }

    // Perform the update operation
    filter := bson.M{"_id": objID}
    _, err = collection.UpdateOne(context.Background(), filter, updateData)
    if err != nil {
        return nil, err
    }

    // Fetch the updated organization
    updatedOrganization, err := GetOrganizationByIDWithMembers(organizationID)
    if err != nil {
        return nil, err
    }

    return updatedOrganization, nil
}

// DeleteOrganizationByID deletes an organization by its ID from the database
func DeleteOrganizationByID(organizationID string) error {
    // Connect to MongoDB
    client, err := GetMongoDBClient()
    if err != nil {
        return err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection("organizations")

    // Convert the organization ID string to MongoDB ObjectID
    objID, err := primitive.ObjectIDFromHex(organizationID)
    if err != nil {
        return err
    }

    // Prepare filter to delete organization
    filter := bson.M{"_id": objID}

    // Delete the organization
    _, err = collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return err
    }

    return nil
}

// SaveInvitation saves an invitation to the database
func SaveInvitation(invitation models.Invitation) error {
    // Connect to MongoDB
    client, err := GetMongoDBClient()
    if err != nil {
        return err
    }
    defer client.Disconnect(context.Background())

    // Get the database and collection
    database := client.Database(DatabaseName)
    collection := database.Collection("invitations")

    // Insert the invitation into the collection
    _, err = collection.InsertOne(context.Background(), invitation)
    if err != nil {
        return err
    }

    return nil
}