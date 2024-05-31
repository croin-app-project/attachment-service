package iservices

import (
	"mime/multipart"

	"github.com/croin-app-project/attachment-service/internal/domain"
)

type IAttachmentService interface {
	CreateAttachment(attachment domain.Attachment, files []*multipart.FileHeader) error
	FindAttachment(systemId, templateId, ref string) ([]domain.File, error)
	DeleteAttachment(systemId, templateId, ref string) error
	UpdateAttachment(systemId, templateId, ref string, files []*multipart.FileHeader) error
}
