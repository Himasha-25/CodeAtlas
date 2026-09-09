package dependencygraph

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) GetGraph(c *gin.Context) {
	runID, _ := strconv.Atoi(c.Query("runId"))
	edges, err := h.svc.GetGraph(uint(runID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, edges)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.GET("/graph", h.GetGraph)
}
