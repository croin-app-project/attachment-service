package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/croin-app-project/attachment-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AttachmentRepository represents a repository interface for managing attachments.
type AttachmentRepository struct {
	collection *mongo.Collection
}

// NewAttachmentRepository creates a new instance of AttachmentRepository.
func NewMongoAttachmentRepository(client *mongo.Client, databaseName, collectionName string) domain.IAttachmentRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &AttachmentRepository{collection: collection}
}

// ExistAsync checks if an attachment exists asynchronously.
func (r *AttachmentRepository) Exist(systemID string, templateID string, ref string) (bool, error) {
	filter := bson.M{"systemId": systemID, "templateId": templateID, "ref": ref}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// Find retrieves an attachment from the repository.
func (r *AttachmentRepository) Find(systemID string, templateID string, ref string) (*domain.Attachment, error) {
	filter := bson.M{"systemId": systemID, "templateId": templateID, "ref": ref}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var attachment domain.Attachment
	err := r.collection.FindOne(ctx, filter).Decode(&attachment)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("attachment not found")
	} else if err != nil {
		return nil, err
	}

	return &attachment, nil
}

// Create creates a new attachment in the repository.
func (r *AttachmentRepository) Create(attachment *domain.Attachment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	attachment.ID = primitive.NewObjectID()

	now := time.Now().UTC()
	attachment.CreatedAt = &now

	_, err := r.collection.InsertOne(ctx, attachment)
	return err
}

// Update updates an existing attachment in the repository.
func (r *AttachmentRepository) Update(attachment *domain.Attachment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	attachment.UpdatedAt = &now

	filter := bson.M{"_id": attachment.ID}
	update := bson.M{"$set": attachment}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

// Delete deletes an attachment from the repository.
func (r *AttachmentRepository) Delete(systemID string, templateID string, ref string) error {
	filter := bson.M{"systemId": systemID, "templateId": templateID, "ref": ref}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
