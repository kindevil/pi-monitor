package route

import "github.com/gin-gonic/gin"

//Server Server
type Server struct {
	Gin  *gin.Engine
	Port string
}

var server *Server

func init() {
	server = &Server{
		Gin:  gin.Default(),
		Port: ":4000",
	}
	route()
}

//Run Run
func Run() {
	server.Gin.Run(server.Port)
}
