package route

import (
	"pi-monitor/service"
)

func route() {
	server.Gin.GET("/", service.GetCPU)
}
