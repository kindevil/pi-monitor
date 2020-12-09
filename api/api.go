/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:00:42
 * @FilePath: /pi-monitor/api/api.go
 * @Date: 2020-12-08 13:37:35
 * @Software: VS Code
 */
package api

import (
	"net/http"
	"pi-monitor/service"

	"github.com/gin-gonic/gin"
)

type statistics struct {
	Host   *service.Host
	CPU    *service.CPU
	Memory *service.Memory
	Net    *service.Net
}

func Collect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Host":   service.GetHost(),
		"CPU":    service.GetCPU(),
		"Memory": service.GetMem(),
		"Net":    service.GetNet(),
	})
}
