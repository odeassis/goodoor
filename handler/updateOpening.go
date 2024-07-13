package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeassis/goodoor/schemas"
)

func UpdateOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusNotFound, fieldfValidate("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s is not faund", id))
		return
	}

	if request.Title != "" {
		opening.Title = request.Title
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error on update opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
