package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "GET openings",
	})
}
