package usecases

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/croin-app-project/attachment-service/internal/domain"
	"github.com/croin-app-project/attachment-service/internal/usecases/iservices"
	_helpers "github.com/croin-app-project/attachment-service/pkg/utils/helpers"
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

func (impl *AttachmentServiceImpl) CreateAttachment(attachment domain.Attachment, files []*multipart.FileHeader) error {
	exist, err := impl._attachmentRepository.Exist(attachment.SystemId, attachment.TemplateId, attachment.Ref)
	if err != nil {
		return err
	} else if exist {
		return errors.New("attachment already exists")
	}

	for index, file := range files {
		fmt.Println(index)
		if file == nil {
			return errors.New("file is required")
		}
		path, err := impl._fileRepository.Save(*file)
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

func (impl *AttachmentServiceImpl) FindAttachment(systemId, templateId, ref string) ([]domain.File, error) {
	attachment, err := impl._attachmentRepository.Find(systemId, templateId, ref)
	if err != nil {
		return nil, err
	}

	paths := _helpers.Map(attachment.Paths, func(path domain.AttachmentPath) string {
		return path.Path
	})

	return impl._fileRepository.GetFiles(paths)
}

func (impl *AttachmentServiceImpl) DeleteAttachment(systemId, templateId, ref string) error {
	attachment, err := impl._attachmentRepository.Find(systemId, templateId, ref)
	if err != nil {
		return err
	}

	paths := _helpers.Map(attachment.Paths, func(path domain.AttachmentPath) string {
		return path.Path
	})

	if err := impl._fileRepository.DeleteFiles(paths); err != nil {
		return err
	}

	return impl._attachmentRepository.Delete(systemId, templateId, ref)
}

func (impl *AttachmentServiceImpl) UpdateAttachment(systemId, templateId, ref string, files []*multipart.FileHeader) error {
	attachment, err := impl._attachmentRepository.Find(systemId, templateId, ref)
	if err != nil {
		return err
	}

	paths := _helpers.Map(attachment.Paths, func(path domain.AttachmentPath) string {
		return path.Path
	})

	if err := impl._fileRepository.DeleteFiles(paths); err != nil {
		return err
	}
	attachment.Paths = []domain.AttachmentPath{}
	for index, file := range files {
		fmt.Println(index)
		if file == nil {
			return errors.New("file is required")
		}
		path, err := impl._fileRepository.Save(*file)
		if err != nil {
			return err
		}
		attachment.Paths = append(attachment.Paths, domain.AttachmentPath{
			No:   index + 1,
			Path: path,
		})
	}

	return impl._attachmentRepository.Update(attachment)
}
