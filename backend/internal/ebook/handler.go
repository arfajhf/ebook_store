package ebook

import (
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
