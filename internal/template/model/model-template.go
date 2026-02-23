package model

import "fmt"

// GetRoutesContent is exported (starts with capital G)
func GetModelContent(entity, nameLower string) string {
	return fmt.Sprintf(`
package model

import (
	"database/sql"
	"time"
)

type %[2]s struct {
	ID          string
	name        sql.NullString
	description sql.NullString
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
`, nameLower, entity)
}
