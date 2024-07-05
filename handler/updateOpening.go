package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpening(cxt *gin.Context) {
	cxt.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
