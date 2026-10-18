package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/go-cicd/internal/todo"
)

var Version = "dev"

func New(store todo.Store) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": Version})
	})
	r.GET("/todos", func(c *gin.Context) {
		items, err := store.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusOK, items)
	})
	r.POST("/todos", func(c *gin.Context) {
		var in struct {
			Title string `json:"title"`
		}
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}
		t, err := store.Create(in.Title)
		if errors.Is(err, todo.ErrEmptyTitle) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusCreated, t)
	})
	return r
}
