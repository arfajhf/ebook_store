package auth

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

func (handler *Handler) Register(context *gin.Context) {
	var request RegisterRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "Data registrasi tidak valid",
		})
		return
	}

	user, err := handler.service.Register(
		context.Request.Context(),
		request,
	)
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailAlreadyUsed):
			context.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "Email sudah digunakan",
			})

		case errors.Is(err, ErrInvalidName):
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  "error",
				"message": "Nama minimal terdiri dari 2 karakter",
			})

		default:
			log.Printf("register user: %v", err)

			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Gagal membuat akun",
			})
		}

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Akun berhasil dibuat",
		"data":    NewUserResponse(user),
	})
}
