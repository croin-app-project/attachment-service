package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	})
}

// func (m *GoMiddleware) ExceptionMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				// Handle the exception
// 				fmt.Println("Recovered from a panic:", r)

// 				// You can log the error, respond with an error message, or perform other actions here

// 				// Respond with an error message
// 				c.JSON(http.StatusInternalServerError, http_response.ErrorResponse{
// 					Message: "An internal server error occurred",
// 					Error:   r,
// 				})

// 				// Abort the request
// 				c.Abort()
// 			}
// 		}()
// 		c.Next()
// 	}
// }

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
