/*
@Time : 2020/12/17 10:56 上午
@Author : jia
@File : cpu_test.go.go
@Software : GoLand
*/

package device

import (
	"fmt"
	"testing"
)

func TestGetCpu(t *testing.T) {
	fmt.Println(GetCpu())
}
