/*
@Time : 2020/12/17 11:20 上午
@Author : jia
@File : disk_test.go
@Software : GoLand
*/

package device

import (
	"fmt"
	"testing"
)

func TestGetDisk(t *testing.T) {
	fmt.Println(GetDisk())
}
