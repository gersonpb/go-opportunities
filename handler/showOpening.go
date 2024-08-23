package handler

import (
	"net/http"

	"github.com/gersonpb/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	
	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening with id: %s not found")
		return
	}
	
	sendSuccess(ctx, "show-opening", opening)

}
