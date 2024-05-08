package usecases

import (
	"errors"
	"fmt"

	"github.com/croin-app-project/attachment-service/internal/adapters/dto"
	"github.com/croin-app-project/attachment-service/internal/domain"
	"github.com/croin-app-project/attachment-service/internal/usecases/iservices"
)

type AttachmentServiceImpl struct {
	_attachmentRepository domain.IAttachmentRepository
	_fileRepository       domain.IFileRepository
}

func NewAttachmentService(attachmentRepository domain.IAttachmentRepository, fileRepository domain.IFileRepository) iservices.IAttachmentService {
	return &AttachmentServiceImpl{
		_attachmentRepository: attachmentRepository,
		_fileRepository:       fileRepository,
	}
}

func (impl *AttachmentServiceImpl) CreateAttachment(req *dto.AttachmentDto) error {
	exist, err := impl._attachmentRepository.Exist(req.SystemId, req.TemplateId, req.Ref)
	if err != nil {
		return err
	} else if exist {
		return errors.New("attachment already exists")
	}

	attachment := &domain.Attachment{
		SystemId:   req.SystemId,
		TemplateId: req.TemplateId,
		Ref:        req.Ref,
	}

	for index, file := range req.Files {
		fmt.Println(index)
		if file == nil {
			return errors.New("file is required")
		}
		path, err := impl._fileRepository.Save(*file) // Pass the value of 'file' instead of its pointer
		if err != nil {
			return err
		}
		attachment.Paths = append(attachment.Paths, domain.AttachmentPath{
			No:   index + 1,
			Path: path,
		})
	}

	return impl._attachmentRepository.Create(attachment)
}
