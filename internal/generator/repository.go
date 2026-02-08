package generator

import (
	"fmt"
	"strings"
)

// GenerateRepository creates the folder and the file for the Repository layer.
func GenerateRepository(entity string) error {
	subFolder := "repository"
	fileName := fmt.Sprintf("%sRepository.go", entity)
	// 2. Define the Template
	// We use a "Raw String Literal" (backticks) to make the Go code readable.
	content := fmt.Sprintf(`package %s

import (
	"context"
)

type %sRepository interface {
	GetByID(ctx context.Context, id string) (interface{}, error)
}

type %sRepositoryImpl struct {
	// Add your DB connection here, e.g., DB *sql.DB
}

func New%sRepository() %sRepository {
	return &%sRepositoryImpl{}
}

func (r *%sRepositoryImpl) GetByID(ctx context.Context, id string) (interface{}, error) {
	return nil, nil
}
`, strings.ToLower(subFolder), entity, entity, entity, entity, entity, entity)

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return SaveGeneratedFile(subFolder, fileName, content)
}
