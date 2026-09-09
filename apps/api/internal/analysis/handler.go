package analysis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) Start(c *gin.Context) {
	repoID, _ := strconv.Atoi(c.Param("repositoryId"))
	run, err := h.svc.Start(uint(repoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, run)
}

func (h *Handler) Status(c *gin.Context) {
	repoID, _ := strconv.Atoi(c.Param("repositoryId"))
	run, err := h.svc.Status(uint(repoID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no analysis found"})
		return
	}
	c.JSON(http.StatusOK, run)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.POST("/repositories/:repositoryId/analysis", h.Start)
	rg.GET("/repositories/:repositoryId/analysis", h.Status)
}
