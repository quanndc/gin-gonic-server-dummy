package apis

import (
	"go_dummy/main/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthyApi(endpoint string, server *core.Server) {
	apis := server.Gin.Group(endpoint)
	// apis.Use() add validator later
	// server.Gin.Use()

	apis.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

}
