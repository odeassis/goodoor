package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeassis/goodoor/schemas"
)

// @BasePath /api/v1
//
// @Description Show a opening job by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpening(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s is not found", id))
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error show opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
