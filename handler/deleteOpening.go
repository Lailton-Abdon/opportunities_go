package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
