package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Ebook Store API is running",
	})
}
