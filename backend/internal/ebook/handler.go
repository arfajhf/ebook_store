package ebook

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) GetAll(context *gin.Context) {
	ebooks, err := handler.service.GetAll(context.Request.Context())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get ebooks",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   ebooks,
	})
}

func (handler *Handler) GetBySlug(context *gin.Context) {
	slug := context.Param("slug")

	item, err := handler.service.GetBySlug(
		context.Request.Context(),
		slug,
	)
	if err != nil {
		log.Printf("get ebook by slug %q: %v", slug, err)

		if errors.Is(err, ErrNotFound) {
			context.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Ebook tidak ditemukan",
			})
			return
		}

		// context.JSON(http.StatusInternalServerError, gin.H{
		// 	"status":  "error",
		// 	"message": "Gagal mengambil detail ebook",
		// 	"debug":   err.Error(),
		// })
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal mengambil detail ebook",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   item,
	})
}
