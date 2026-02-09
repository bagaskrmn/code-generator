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

	files := map[string]string{
		fmt.Sprintf("%s.go", nameLower): fmt.Sprintf(`%s.go`, entity),
	}
	for fileName, content := range files {
		fullPath := filepath.Join(targetDir, fileName)
		err := os.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", fileName, err)
		}
		fmt.Printf("Created: %s\n", fullPath)
	}

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return nil
}
