/*
@Time : 2020/12/17 3:59 下午
@Author : jia
@File : float.go.go
@Software : GoLand
*/

package helper

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ToFixed(value float64) float64 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	if err != nil {
		log.Error(err)
	}
	return value
}
