package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpenings_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
