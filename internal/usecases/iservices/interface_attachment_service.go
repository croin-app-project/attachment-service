package iservices

import "github.com/croin-app-project/attachment-service/internal/adapters/dto"

type IAttachmentService interface {
	CreateAttachment(attachment *dto.AttachmentDto) error
}
