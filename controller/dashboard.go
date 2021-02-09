/*
@Time : 2020/12/17 11:23 上午
@Author : jia
@File : dashboard.go
@Software : GoLand
*/

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
