/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:21:23
 * @FilePath: /pi-monitor/route/route.go
 * @Date: 2020-12-06 19:58:05
 * @Software: VS Code
 */
package route

import (
	"pi-monitor/api"
	"pi-monitor/controller"
)

func route() {
	server.Gin.LoadHTMLGlob("views/**/*")
	server.Gin.Static("static/", "static/")
	server.Gin.GET("/", controller.Dashboard)
	server.Gin.GET("/api/get", api.Collect)
}
