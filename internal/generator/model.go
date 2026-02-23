package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bagaskrmn/code-generator/internal/template/model"
)

func GenerateModel(entity string) error {
	nameLower := strings.ToLower(entity)
	targetDir := GetPath("model", entity)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create model directory: %w", err)
	}
	fileName := fmt.Sprintf("%s.go", nameLower)
	fullPath := filepath.Join(targetDir, fileName)

	// content
	content := model.GetModelContent(entity, nameLower)

	// write file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write model file: %w", err)
	}

	fmt.Printf("Created: %s\n", fullPath)
	return nil

}
