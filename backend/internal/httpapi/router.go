package httpapi

import (
	"ebook-store/backend/internal/ebook"
	"ebook-store/backend/internal/health"

	"github.com/gin-gonic/gin"
)

func NewRouter(ebookHandler *ebook.Handler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/health", health.Check)
		api.GET("/ebooks", ebookHandler.GetAll)
		api.GET("/ebooks/:slug", ebookHandler.GetBySlug)
	}

	return router
}
