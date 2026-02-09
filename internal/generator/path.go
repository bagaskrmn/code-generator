package generator

import (
	"path/filepath"
	"strings"
	"os"
)

// GetPath returns the target directory for a specific module and entity.
func GetPath(module, entity string) string {
	nameLower := strings.ToLower(entity)
	// 1. Determine the Base Path
	base := ""
	// Check if we should generate into "generated-code" or root
	if strings.ToLower(os.Getenv("AUTO_IMPLEMENT")) != "true" {
		base = "generated-code"
	}

	switch strings.ToLower(module) {
	case "repository":
		return filepath.Join(base, "pkg", nameLower, "repository")
	case "service":
		return filepath.Join(base, "pkg", nameLower, "service")
	case "handler":
		return filepath.Join(base, "api", "handlers", nameLower)
	case "presenter":
		return filepath.Join(base, "api", "presenter", nameLower)
	case "routes":
		// Routes are grouped together in one folder usually
		return filepath.Join(base, "api", "routes")
	default:
		return base
	}
}
