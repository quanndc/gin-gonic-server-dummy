package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin  *gin.Engine
	Port string
}

func NewServer() *Server {
	server := &Server{
		Port: "8080",
		Gin:  gin.Default(),
	}

	return server
}

func (server *Server) Start() error {
	server.Gin.SetTrustedProxies(nil)
	server.Gin.Use(cors.Default())

	return server.Gin.Run(":" + server.Port)
}
