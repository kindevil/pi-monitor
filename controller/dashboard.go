/*
@Time : 2020/12/17 11:23 上午
@Author : jia
@File : dashboard.go
@Software : GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pi-monitor/device"
)

func Dashboard(c *gin.Context) {
	var data = gin.H{}
	var host = device.GetHost()

	data["title"] = host.Vendor + " Monitor"
	data["host"] = host

	c.HTML(http.StatusOK, "dashboard.tmpl", data)
}
