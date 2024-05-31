package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attachment struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedBy  *string            `bson:"createdBy"`
	CreatedAt  *time.Time         `bson:"createdAt"`
	UpdatedBy  *string            `bson:"updatedBy"`
	UpdatedAt  *time.Time         `bson:"updatedAt"`
	SystemId   string             `bson:"systemId"`
	TemplateId string             `bson:"templateId"`
	Ref        string             `bson:"ref"`
	Paths      []AttachmentPath   `bson:"paths"`
}

type AttachmentPath struct {
	No   int    `json:"no"`
	Path string `json:"path"`
}

type IAttachmentRepository interface {
	Exist(systemId string, templateId string, ref string) (bool, error)
	Find(systemId string, templateId string, ref string) (*Attachment, error)
	Create(attachment Attachment) error
	Update(attachment *Attachment) error
	Delete(systemId string, templateId string, ref string) error
}
