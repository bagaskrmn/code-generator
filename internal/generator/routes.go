package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bagaskrmn/code-generator/internal/template/routes"
)

func GenerateRoutes(entity string, repoName string) error {
	nameLower := strings.ToLower(entity)
	targetDir := GetPath("routes", entity)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create routes directory: %w", err)
	}
	fileName := fmt.Sprintf("%s.go", nameLower)
	fullPath := filepath.Join(targetDir, fileName)

	// content
	content := routes.GetRoutesContent(entity, nameLower, repoName)

	// write file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write routes file: %w", err)
	}

	fmt.Printf("Created: %s\n", fullPath)
	return nil

}
