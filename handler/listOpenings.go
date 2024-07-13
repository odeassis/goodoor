package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeassis/goodoor/schemas"
)

func ListOpenings(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "an error in list opening")
		return
	}

	sendSuccess(ctx, "list-openings", opening)
}
