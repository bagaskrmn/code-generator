package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// Exported helper to handle all disk operations
func SaveGeneratedFile(subFolder, fileName, content string) error {
	// 1. Define the root output directory
	baseDir := "generated-code"
	
	// 2. Build the full path: generated-code/handler/UserHandler.go
	targetPath := filepath.Join(baseDir, subFolder)
	fullFilePath := filepath.Join(targetPath, fileName)

	// 3. Create the directory tree (0755 = standard permissions)
	// This will create 'generated-code' and 'subFolder' automatically
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
	}

	// 4. Write the file to disk
	if err := os.WriteFile(fullFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fullFilePath, err)
	}

	fmt.Printf("  ✅ Created: %s\n", fullFilePath)
	return nil
}