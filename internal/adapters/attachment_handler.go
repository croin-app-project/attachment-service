package adapters

import (
	"github.com/croin-app-project/attachment-service/internal/adapters/dto"
	"github.com/croin-app-project/attachment-service/internal/domain"
	"github.com/croin-app-project/attachment-service/internal/usecases/iservices"
	http_response "github.com/croin-app-project/attachment-service/pkg/utils/http-response"
	"github.com/gofiber/fiber/v2"
)

type AttachmentControllerImpl struct {
	_attachmentService iservices.IAttachmentService
}

func NewAttachmentController(f fiber.Router, attachmentService iservices.IAttachmentService) {
	handler := &AttachmentControllerImpl{_attachmentService: attachmentService}
	v1 := f.Group("/v1")
	api := v1.Group("/attachment")
	api.Post("", handler.Create)
	api.Get("", handler.Find)
	api.Delete("", handler.Delete)
	api.Put("", handler.Update)
}

func (impl *AttachmentControllerImpl) Create(c *fiber.Ctx) error {
	var body dto.AttachmentCreateDto
	if err := c.QueryParser(&body); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INVALID_INPUT_PARAMETER, err)
		return c.Status(errCode).JSON(errObj)
	}

	form, _ := c.MultipartForm()

	files := form.File["files"]

	attachment := domain.Attachment{
		SystemId:   body.SystemId,
		TemplateId: body.TemplateId,
		Ref:        body.Ref,
	}

	if err := impl._attachmentService.CreateAttachment(attachment, files); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INTERNAL_SYSTEM_ERROR, err)
		return c.Status(errCode).JSON(errObj)
	} else {
		return c.Status(fiber.StatusCreated).JSON(http_response.SuccessReponse{
			Code:   fiber.StatusCreated,
			Status: "success!",
			Result: dto.AttachmentCreatePresenter{
				Message: "Attachment created successfully!",
			},
		})
	}
}

func (impl *AttachmentControllerImpl) Find(c *fiber.Ctx) error {
	body := dto.AttachmentFindDto{}
	if err := c.QueryParser(&body); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INVALID_INPUT_PARAMETER, err)
		return c.Status(errCode).JSON(errObj)
	}

	files, err := impl._attachmentService.FindAttachment(body.SystemId, body.TemplateId, body.Ref)
	if err != nil {
		errCode, errObj := http_response.HandleException(http_response.INTERNAL_SYSTEM_ERROR, err)
		return c.Status(errCode).JSON(errObj)
	} else {
		return c.Status(fiber.StatusOK).JSON(http_response.SuccessReponse{
			Code:   fiber.StatusOK,
			Status: "success!",
			Result: dto.AttachmentFindPresenter{
				Files: files,
			},
		})
	}
}

func (impl *AttachmentControllerImpl) Delete(c *fiber.Ctx) error {
	body := dto.AttachmentDeleteDto{}
	if err := c.QueryParser(&body); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INVALID_INPUT_PARAMETER, err)
		return c.Status(errCode).JSON(errObj)
	}

	if err := impl._attachmentService.DeleteAttachment(body.SystemId, body.TemplateId, body.Ref); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INTERNAL_SYSTEM_ERROR, err)
		return c.Status(errCode).JSON(errObj)
	} else {
		return c.Status(fiber.StatusOK).JSON(http_response.SuccessReponse{
			Code:   fiber.StatusOK,
			Status: "success!",
			Result: dto.AttachmentDeletePresenter{
				Message: "Attachment deleted successfully!",
			},
		})
	}
}

func (impl *AttachmentControllerImpl) Update(c *fiber.Ctx) error {
	body := dto.AttachmentUpdateDto{}
	if err := c.QueryParser(&body); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INVALID_INPUT_PARAMETER, err)
		return c.Status(errCode).JSON(errObj)
	}

	form, _ := c.MultipartForm()

	files := form.File["files"]

	if err := impl._attachmentService.UpdateAttachment(body.SystemId, body.TemplateId, body.Ref, files); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INTERNAL_SYSTEM_ERROR, err)
		return c.Status(errCode).JSON(errObj)
	} else {
		return c.Status(fiber.StatusOK).JSON(http_response.SuccessReponse{
			Code:   fiber.StatusOK,
			Status: "success!",
			Result: dto.AttachmentUpdatePresenter{
				Message: "Attachment updated successfully!",
			},
		})
	}

}
