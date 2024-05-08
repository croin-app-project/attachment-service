package adapters

import (
	"net/http"

	"github.com/croin-app-project/attachment-service/internal/adapters/dto"
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
	api.Post("/create", handler.Create)
}

func (impl *AttachmentControllerImpl) Create(c *fiber.Ctx) error {
	var body dto.AttachmentDto

	if err := c.BodyParser(&body); err != nil {
		errCode, errObj := http_response.HandleException(http_response.INVALID_INPUT_PARAMETER, err)
		return c.Status(errCode).JSON(errObj)
	}
	form, _ := c.MultipartForm()
	files := form.File["Files"]

	body.Files = files
	err := impl._attachmentService.CreateAttachment(&body)
	if err != nil {
		errCode, errObj := http_response.HandleException(http_response.INTERNAL_SYSTEM_ERROR, err)
		return c.Status(errCode).JSON(errObj)
	}

	result := dto.AttachmentCreatePresenter{
		Message: "Attachment created successfully!",
	}

	errCode, errObj := http.StatusOK, http_response.SuccessReponse{Code: http.StatusOK,
		Status: "success!",
		Result: result,
	}
	return c.Status(errCode).JSON(errObj)
}
