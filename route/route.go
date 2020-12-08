/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:01:09
 * @FilePath: /pi-monitor/route/route.go
 * @Date: 2020-12-06 19:58:05
 * @Software: VS Code
 */
package route

import (
	"net/http"
	"pi-monitor/api"

	"github.com/gin-gonic/gin"
)

func route() {
	server.Gin.LoadHTMLGlob("views/**/*")
	server.Gin.Static("static/", "static/")

	data := gin.H{}
	data["title"] = "Raspberrypi Monitor"

	server.Gin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", data)
	})

	server.Gin.GET("/api/get", api.Collect)
}
