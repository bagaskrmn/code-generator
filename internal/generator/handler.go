package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateHandler(entity string) error {
	nameLower := strings.ToLower(entity)

	targetDir := GetPath("handler", entity)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create handler directory: %w", err)
	}

	files := map[string]string{
		fmt.Sprintf("create_%s_handlers.go", nameLower): fmt.Sprintf(`create_%s_handlers.go`, entity),
		fmt.Sprintf("read_%s_handlers.go", nameLower):   fmt.Sprintf(`read_%s_handlers.go`, entity),
		fmt.Sprintf("update_%s_handlers.go", nameLower): fmt.Sprintf(`update_%s_handlers.go`, entity),
		fmt.Sprintf("delete_%s_handlers.go", nameLower): fmt.Sprintf(`delete_%s_handlers.go`, entity),
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
