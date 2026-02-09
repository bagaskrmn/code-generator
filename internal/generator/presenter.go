package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GeneratePresenter(entity string) error {
	nameLower := strings.ToLower(entity)

	targetDir := filepath.Join("generated-code", "api", "presenter", nameLower)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create presenter directory: %w", err)
	}

	files := map[string]string{
		fmt.Sprintf("req_create_%s.go", nameLower):    fmt.Sprintf(`req_create_%s.go`, entity),
		fmt.Sprintf("req_update_%s.go", nameLower):    fmt.Sprintf(`req_update_%s.go`, entity),
		fmt.Sprintf("req_delete_%s.go.go", nameLower): fmt.Sprintf(`req_delete_%s.go`, entity),
		fmt.Sprintf("req_read_%s.go.go", nameLower):   fmt.Sprintf(`req_read_%s.go`, entity),
		fmt.Sprintf("res_create_%s.go", nameLower):    fmt.Sprintf(`res_create_%s.go`, entity),
		fmt.Sprintf("res_update_%s.go", nameLower):    fmt.Sprintf(`res_update_%s.go`, entity),
		fmt.Sprintf("res_delete_%s.go.go", nameLower): fmt.Sprintf(`res_delete_%s.go`, entity),
		fmt.Sprintf("res_read_%s.go.go", nameLower):   fmt.Sprintf(`res_read_%s.go`, entity),
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
