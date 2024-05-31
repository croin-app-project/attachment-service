package dto

import (
	"mime/multipart"

	"github.com/croin-app-project/attachment-service/internal/domain"
)

type AttachmentCreateDto struct {
	SystemId   string                  `form:"systemId" query:"systemId" json:"systemId" binding:"required"`
	TemplateId string                  `form:"templateId" query:"templateId" json:"templateId" binding:"required"`
	Ref        string                  `form:"ref" query:"ref" json:"ref" binding:"required"`
	Files      []*multipart.FileHeader `form:"files" binding:"required"`
}
type AttachmentCreatePresenter struct {
	Message string `json:"message"`
}

type AttachmentFindDto struct {
	SystemId   string `form:"systemId" query:"systemId" json:"systemId" binding:"required"`
	TemplateId string `form:"templateId" query:"templateId" json:"templateId" binding:"required"`
	Ref        string `form:"ref" query:"ref" json:"ref" binding:"required"`
}
type AttachmentFindPresenter struct {
	Files []domain.File `bson:"paths"`
}

type AttachmentDeleteDto struct {
	SystemId   string `form:"systemId" query:"systemId" json:"systemId" binding:"required"`
	TemplateId string `form:"templateId" query:"templateId" json:"templateId" binding:"required"`
	Ref        string `form:"ref" query:"ref" json:"ref" binding:"required"`
}

type AttachmentDeletePresenter struct {
	Message string `json:"message"`
}

type AttachmentUpdateDto struct {
	SystemId   string                  `form:"systemId"  query:"systemId"  json:"systemId" binding:"required"`
	TemplateId string                  `form:"templateId"  query:"templateId"  json:"templateId" binding:"required"`
	Ref        string                  `form:"ref"  query:"ref"  json:"ref" binding:"required"`
	Files      []*multipart.FileHeader `form:"files"`
}

type AttachmentUpdatePresenter struct {
	Message string `json:"message"`
}
