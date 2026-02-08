package generator

import (
	"fmt"
	"strings"
)

func GenerateHandler(entity string) error {
	subFolder := "handler"
	fileName := fmt.Sprintf("%sHandler.go", entity)

	content := fmt.Sprintf(`package %s

import "github.com/gofiber/fiber/v2"

type %sHandler struct {
	// Inject Service here
}

func New%sHandler() *%sHandler {
	return &%sHandler{}
}

func (h *%sHandler) GetData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello %s"})
}
`, strings.ToLower(subFolder), entity, entity, entity, entity, entity, entity)

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return SaveGeneratedFile(subFolder, fileName, content)
}