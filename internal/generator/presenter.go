package generator

import (
	"fmt"
	"strings"
)

func GeneratePresenter(entity string) error {
	subFolder := "presenter"
	fileName := fmt.Sprintf("%sPresenter.go", entity)

	content := fmt.Sprintf(`package %s

type %sResponse struct {
	ID   string ` + "`" + `json:"id"` + "`" + `
	Name string ` + "`" + `json:"name"` + "`" + `
}

func To%sResponse(data interface{}) *%sResponse {
	return &%sResponse{}
}
`, strings.ToLower(subFolder), entity, entity, entity, entity)

	// Call the centralized writer instead of os.MkdirAll and os.WriteFile
	return SaveGeneratedFile(subFolder, fileName, content)
}