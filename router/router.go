package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initialize_Routes(r)

	r.Run(":8080")
}
