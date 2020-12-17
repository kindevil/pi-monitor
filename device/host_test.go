/*
@Time : 2020/12/17 10:28 上午
@Author : jia
@File : host_test.go.go
@Software : GoLand
*/

package device

import (
	"fmt"
	"testing"
)

func TestGetHost(t *testing.T) {
	fmt.Println(GetHost())
}
