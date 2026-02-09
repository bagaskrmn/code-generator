package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateRepository(entity string) error {
	nameLower := strings.ToLower(entity)

	// 1. Define target directory: root/pkg/<name>/repository/
	targetDir := filepath.Join("generated-code", "pkg", nameLower, "repository")

	// 2. Create the directory tree
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create repository directory: %w", err)
	}

	// 3. Define the blueprint for the 5 files
	files := map[string]string{
		// The main interface and struct definition
		"repository.go": fmt.Sprintf(`%s repository.go`, entity),
	}

	// 4. Write each file independently
	for fileName, content := range files {
		fullPath := filepath.Join(targetDir, fileName)
		err := os.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", fileName, err)
		}
		fmt.Printf("Created: %s\n", fullPath)
	}

	return nil
}
