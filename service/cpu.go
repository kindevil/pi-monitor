package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/wonderivan/logger"
)

func GetCPU(c *gin.Context) {
	info, err := cpu.Info()
	if err != nil {
		logger.Error(err)
	}
	c.String(http.StatusOK, "%s", info)
}
