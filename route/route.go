/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:21:23
 * @FilePath: /pi-monitor/route/route.go
 * @Date: 2020-12-06 19:58:05
 * @Software: VS Code
 */
package route

import (
	"pi-monitor/service"
	"pi-monitor/websocket"
)

func route() {
	server.Gin.LoadHTMLFiles("web/index.html")
	server.Gin.Static("assets/", "web/assets/")
	server.Gin.GET("/", service.Dashboard)
	server.Gin.GET("/ws", websocket.HandleWebSocket)
}
