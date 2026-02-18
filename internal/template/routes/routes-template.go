package routes

import "fmt"

// GetRoutesContent is exported (starts with capital G)
func GetRoutesContent(entity, nameLower string, repoName string) string {
	return fmt.Sprintf(`

package routes

import (
	%[1]sHandlers "github.com/ABA-Developer/%[3]s/api/handlers/%[1]s"
	"github.com/ABA-Developer/%[3]s/api/middleware"
	"github.com/ABA-Developer/%[3]s/internal/module"
	%[1]sSvc "github.com/ABA-Developer/%[3]s/pkg/%[1]s/service"

	"github.com/gofiber/fiber/v2"
)

func add%[2]sRoutes(f fiber.Router, m *module.Module) {
	%[2]sSvc := %[1]sSvc.NewService(m.GetDB(), m.GetLogger())
	mdw := middleware.NewEnsureToken(m.GetDB())

	%[2]sGroup := f.Group("/%[1]s", mdw.ValidateToken(), middleware.RateLimiterMiddleware(m.GetRuntimeConfig(), true))

	%[2]sGroup.Get("", %[2]sHandlers.ReadList%[2]sApp(%[1]sSvc, m.GetLoggerFile()))
	%[2]sGroup.Get("/:id", %[2]sHandlers.ReadDetail%[2]sApp(%[1]sSvc, m.GetLoggerFile()))
	%[2]sGroup.Post("", %[2]sHandlers.Create%[2]sApp(%[1]sSvc, m.GetValidator(), m.GetLoggerFile()))
	%[2]sGroup.Put("/:id", %[2]sHandlers.Update%[2]sApp(%[1]sSvc, m.GetValidator(), m.GetLoggerFile()))
	%[2]sGroup.Delete("/:id", %[2]sHandlers.Delete%[2]sApp(%[1]sSvc, m.GetLoggerFile()))
}

`, nameLower, entity, repoName)
}
