/*
@Time : 2020/12/17 11:12 上午
@Author : jia
@File : mem_test.go.go
@Software : GoLand
*/

package device

import (
	"fmt"
	"testing"
)

func TestGetMemory(t *testing.T) {
	fmt.Println(GetMemory())
}
