package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeassis/goodoor/schemas"
)

// @BaseParth /api/v1
//
// @Description Show all openings jobs
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpenings(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "an error in list opening")
		return
	}

	sendSuccess(ctx, "list-openings", opening)
}
