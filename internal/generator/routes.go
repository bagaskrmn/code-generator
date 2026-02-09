package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateRoutes(entity string) error {
	nameLower := strings.ToLower(entity)

	targetDir := GetPath("routes", entity)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create routes directory: %w", err)
	}
	fileName := fmt.Sprintf("%s.go", nameLower)
	fullPath := filepath.Join(targetDir, fileName)

	content := fmt.Sprintf(`package routes

import "github.com/gofiber/fiber/v2"

func Setup%sRoutes(app *fiber.App) {
	api := app.Group("/%s")
	
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("%s endpoint")
	})
}
`, entity, nameLower, entity)
	// 3. Write directly to disk
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write routes file: %w", err)
	}

	fmt.Printf("Created: %s\n", fullPath)
	return nil

}
