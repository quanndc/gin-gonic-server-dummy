package apis

import (
	"go_dummy/main/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthyApi(endpoint string, server *core.Server) gin.IRoutes {
	apis := server.Gin.Group(endpoint)

	// apis.Use() add validator later

	apis.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Healthy",
		})
	})
	return apis
}
