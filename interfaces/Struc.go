package interfaces

import (
	"go/types"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Cause string
}
type StandarMessage struct {
	Message string
	Data    *types.Struct
}

func ErrorHandling(c *fiber.Ctx, cause Error, status fiber.Error) error {
	return c.Status(status.Code).JSON(cause)
}
