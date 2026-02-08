package generator

import (
	"fmt"
	"strings"
)

func GenerateRoutes(entity string) error {
	subFolder := "routes"
	fileName := fmt.Sprintf("%sRoutes.go", entity)

	content := fmt.Sprintf(`package %s

import "github.com/gofiber/fiber/v2"

func Setup%sRoutes(app *fiber.App) {
	api := app.Group("/%s")
	
	// Define endpoints
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("%s endpoint")
	})
}
`, strings.ToLower(subFolder), entity, strings.ToLower(entity), entity)

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return SaveGeneratedFile(subFolder, fileName, content)
}