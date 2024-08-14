package handler

import (
	"net/http"

	"github.com/Lailton-Abdon/opportunities_go/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpenings_handler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
