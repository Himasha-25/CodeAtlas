package codeexplorer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) ListFiles(c *gin.Context) {
	runID, _ := strconv.Atoi(c.Query("runId"))
	files, err := h.svc.ListFiles(uint(runID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

func (h *Handler) ListSymbols(c *gin.Context) {
	fileID, _ := strconv.Atoi(c.Param("fileId"))
	symbols, err := h.svc.ListSymbols(uint(fileID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, symbols)
}

func (h *Handler) Search(c *gin.Context) {
	runID, _ := strconv.Atoi(c.Query("runId"))
	symbols, err := h.svc.Search(uint(runID), c.Query("q"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, symbols)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.GET("/explorer/files", h.ListFiles)
	rg.GET("/explorer/files/:fileId/symbols", h.ListSymbols)
	rg.GET("/explorer/search", h.Search)
}
