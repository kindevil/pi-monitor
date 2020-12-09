package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	host := GetHost()

	data := gin.H{}
	data["title"] = host.Vendor + " Monitor"
	data["hostname"] = host.Hostname
	data["os"] = host.OS
	data["vendor"] = host.Vendor
	data["model"] = host.Model
	data["serial"] = host.Serial
	data["boottime"] = host.BootTime
	data["kernal"] = host.Kernal
	data["netCount"] = host.InterfaceNum

	c.HTML(http.StatusOK, "home.tmpl", data)
}

// func netjs() {
// 	n := GetNet()

// 	for _, val := range n.Interface {
// 		js := `var net` + val.Name + ` = echarts.init(document.getElementById('net'));`
// 	}

// }
