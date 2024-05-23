package router

import (
	"github.com/Lailton-Abdon/opportunities_go/handler"
	"github.com/gin-gonic/gin"
)

func initialize_Routes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpening_handler)
		v1.POST("/opening", handler.CreateOpening_handler)
		v1.PUT("/opening", handler.UpdateOpening_handler)
		v1.DELETE("/opening", handler.DeleteOpening_handler)
		v1.GET("/openings", handler.ListOpenings_handler)
	}
}
