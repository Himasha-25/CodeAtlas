package httpserver

import (
	"github.com/codeatlas/api/internal/analysis"
	"github.com/codeatlas/api/internal/assistant"
	"github.com/codeatlas/api/internal/auth"
	"github.com/codeatlas/api/internal/codeexplorer"
	"github.com/codeatlas/api/internal/config"
	"github.com/codeatlas/api/internal/dependencygraph"
	docs "github.com/codeatlas/api/internal/documentation"
	"github.com/codeatlas/api/internal/impact"
	"github.com/codeatlas/api/internal/platform/middleware"
	"github.com/codeatlas/api/internal/project"
	"github.com/codeatlas/api/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.CORS())

	v1 := r.Group("/api/v1")

	repoSvc := repository.NewService(db)

	auth.RegisterRoutes(v1, auth.NewService(db, cfg.JWTSecret))

	protected := v1.Group("")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	project.RegisterRoutes(protected, project.NewService(db))
	repository.RegisterRoutes(protected, repoSvc)
	analysis.RegisterRoutes(protected, analysis.NewService(db, repoSvc))
	codeexplorer.RegisterRoutes(protected, codeexplorer.NewService(db))
	dependencygraph.RegisterRoutes(protected, dependencygraph.NewService(db))
	docs.RegisterRoutes(protected, docs.NewService(db))
	assistant.RegisterRoutes(protected, assistant.NewService(db, cfg.GroqAPIKey))
	impact.RegisterRoutes(protected, impact.NewService(db))

	return r
}
