package main

import (
	"context"
	"fmt"

	_config "github.com/croin-app-project/attachment-service/config"
	_middleware "github.com/croin-app-project/attachment-service/middleware"

	"github.com/croin-app-project/attachment-service/internal/adapters"
	_repository "github.com/croin-app-project/attachment-service/internal/domain/repositories"
	_service "github.com/croin-app-project/attachment-service/internal/usecases"
	_helpers "github.com/croin-app-project/package/pkg/utils/helpers"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config := _config.ReadConfiguration()
	dbConfig := _helpers.Filter(config.Databases, func(s _config.DatabaseSetting) bool {
		return s.DbName == "attachment"
	})[0]
	configService := _helpers.Filter(config.Server.Services, func(s _config.ServiceSetting) bool {
		return s.Name == "attachment-service"
	})[0]

	clientOptions := options.Client().ApplyURI(dbConfig.Url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("mongo connect: ", err)
	}
	defer client.Disconnect(context.Background())

	app := fiber.New()

	middL := _middleware.InitMiddleware()
	app.Use(middL.CORS())

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("ok!")
	})
	api := app.Group("/api/" + configService.Name)

	attachmentRepository := _repository.NewMongoAttachmentRepository(client, dbConfig.DbName, (*dbConfig.Collections)["attachments"])
	fileRepository := _repository.NewFileRepository()
	attachmentService := _service.NewAttachmentService(attachmentRepository, fileRepository)
	adapters.NewAttachmentController(api, attachmentService)

	app.Listen(":" + configService.Port)
}
