package documentation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) Get(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("projectId"))
	doc, err := h.svc.Get(uint(projectID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no documentation found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

func (h *Handler) Generate(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("projectId"))
	runID, _ := strconv.Atoi(c.Query("runId"))
	doc, err := h.svc.Generate(uint(projectID), uint(runID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, doc)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.GET("/projects/:projectId/docs", h.Get)
	rg.POST("/projects/:projectId/docs/generate", h.Generate)
}
