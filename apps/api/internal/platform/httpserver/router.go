package httpserver

import (
	"github.com/codeatlas/api/internal/config"
	"github.com/codeatlas/api/internal/platform/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.CORS())

	v1 := r.Group("/api/v1")

	// Auth routes (no middleware)
	// auth.RegisterRoutes(v1, db, cfg)

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	// project.RegisterRoutes(protected, db)
	// repository.RegisterRoutes(protected, db)
	// analysis.RegisterRoutes(protected, db)
	// codeexplorer.RegisterRoutes(protected, db)
	// dependencygraph.RegisterRoutes(protected, db)
	// documentation.RegisterRoutes(protected, db)
	// assistant.RegisterRoutes(protected, db, cfg)
	// impact.RegisterRoutes(protected, db)

	return r
}
