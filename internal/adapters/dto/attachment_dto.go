package dto

import (
	"mime/multipart"

	"github.com/croin-app-project/attachment-service/internal/domain"
)

type AttachmentDto struct {
	SystemId   string                  `form:"systemId" json:"systemId" binding:"required"`
	TemplateId string                  `form:"templateId" json:"templateId" binding:"required"`
	Ref        string                  `form:"ref" json:"ref" binding:"required"`
	Files      []*multipart.FileHeader `form:"files" binding:"required"`
	// File *fiber.FileHeader `form:"file" binding:"required"`
}
type AttachmentCreatePresenter struct {
	Message string `json:"message"`
}

type AttachmentFindPresenter struct {
	Paths []domain.AttachmentPath `bson:"paths"`
}
