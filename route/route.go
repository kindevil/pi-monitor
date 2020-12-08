/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:21:23
 * @FilePath: /pi-monitor/route/route.go
 * @Date: 2020-12-06 19:58:05
 * @Software: VS Code
 */
package route

import (
	"net/http"
	"pi-monitor/api"
	"pi-monitor/service"

	"github.com/gin-gonic/gin"
)

func route() {
	server.Gin.LoadHTMLGlob("views/**/*")
	server.Gin.Static("static/", "static/")

	host := service.GetHost()

	data := gin.H{}
	data["title"] = "Raspberrypi Monitor"
	data["boottime"] = host.BootTime
	data["version"] = host.Version
	data["serial"] = host.Serial
	data["hostname"] = host.Hostname
	data["platform"] = host.Platform
	data["kernal"] = host.Kernal
	data["hardware"] = host.Hardware
	data["model"] = host.Model
	data["os"] = host.OS

	server.Gin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", data)
	})

	server.Gin.GET("/api/get", api.Collect)
}
