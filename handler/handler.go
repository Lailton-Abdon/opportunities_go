package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func CreateOpening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

func UpdateOpening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func DeleteOpening_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func ListOpenings_handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
