package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateService(entity string) error {
	nameLower := strings.ToLower(entity)

	// 1. Define target directory: root/pkg/<name>/service/
	targetDir := filepath.Join("generated-code", "pkg", nameLower, "service")

	// 2. Create the directory tree
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create service directory: %w", err)
	}

	// 3. Define the blueprint for the 5 files
	files := map[string]string{
		// The main interface and struct definition
		"service.go": fmt.Sprintf(`%s service.go`, entity),

		// Implementation: Create
		fmt.Sprintf("create_%s_service.go", nameLower): fmt.Sprintf(`create_%s_service.go`, entity),

		// Implementation: Read
		fmt.Sprintf("read_%s_service.go", nameLower): fmt.Sprintf(`read_%s_service.go`, entity),

		// Implementation: Update
		fmt.Sprintf("update_%s_service.go", nameLower): fmt.Sprintf(`update_%s_service.go`, entity),

		// Implementation: Delete
		fmt.Sprintf("delete_%s_service.go", nameLower): fmt.Sprintf(`delete_%s_service.go`, entity),
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
