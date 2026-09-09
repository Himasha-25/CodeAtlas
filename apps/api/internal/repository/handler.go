package repository

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func NewHandler(svc *Service) *Handler { return &Handler{svc} }

func (h *Handler) Upload(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("projectId"))
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}
	defer file.Close()
	repo, err := h.svc.Upload(uint(projectID), header.Filename, file, header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, repo)
}

func (h *Handler) List(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("projectId"))
	repos, err := h.svc.List(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repos)
}

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
	h := NewHandler(svc)
	rg.POST("/projects/:projectId/repositories", h.Upload)
	rg.GET("/projects/:projectId/repositories", h.List)
}
