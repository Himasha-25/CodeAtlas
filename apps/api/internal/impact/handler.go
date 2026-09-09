package impact

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) Analyze(c *gin.Context) {
	runID, _ := strconv.Atoi(c.Query("runId"))
	symbolID, _ := strconv.Atoi(c.Query("symbolId"))
	result, err := h.svc.Analyze(uint(runID), uint(symbolID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.GET("/impact", h.Analyze)
}
