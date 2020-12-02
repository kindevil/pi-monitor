package route

import "github.com/gin-gonic/gin"

type Server struct {
	Gin  *gin.Engine
	Port string
}

var server *Server

func init() {
	server = &Server{
		Gin:  gin.Default(),
		Port: ":3000",
	}
	route()
}

func Run() {
	server.Gin.Run(server.Port)
}
