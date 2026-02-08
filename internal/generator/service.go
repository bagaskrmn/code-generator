package generator

import (
	"fmt"
	"strings"
)

func GenerateService(entity string) error {
	subFolder := "service"
	fileName := fmt.Sprintf("%sService.go", entity)

	content := fmt.Sprintf(`package %s

type %sService interface {
	Execute(id string) error
}

type %sServiceImpl struct {
	// Inject Repository here
}

func New%sService() %sService {
	return &%sServiceImpl{}
}

func (s *%sServiceImpl) Execute(id string) error {
	// Business logic for %s goes here
	return nil
}
`, strings.ToLower(subFolder), entity, entity, entity, entity, entity, entity, entity)

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return SaveGeneratedFile(subFolder, fileName, content)
}