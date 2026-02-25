package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/bagaskrmn/code-generator/internal/template/handler"
)

func GenerateHandler(entity string) error {
	nameLower := strings.ToLower(entity)

	targetDir := GetPath("handler", entity)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create handler directory: %w", err)
	}

	files := map[string]string{
        fmt.Sprintf("create_%s_handlers.go", nameLower):      handler.GetCreateHandlerContent(entity, nameLower),
        fmt.Sprintf("read_%s_handlers.go", nameLower):        handler.GetReadHandlerContent(entity, nameLower),
        fmt.Sprintf("read_detail_%s_handlers.go", nameLower): handler.GetReadDetailHandlerContent(entity, nameLower),
        fmt.Sprintf("update_%s_handlers.go", nameLower):      handler.GetUpdateHandlerContent(entity, nameLower),
        fmt.Sprintf("delete_%s_handlers.go", nameLower):      handler.GetDeleteHandlerContent(entity, nameLower),
    }
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
