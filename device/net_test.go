/*
@Time : 2020/12/17 10:58 上午
@Author : jia
@File : net_test.go.go
@Software : GoLand
*/

package device

import (
	"fmt"
	"testing"
)

func TestGetNetNames(t *testing.T) {
	fmt.Println(GetNetNames())
}

func TestGetNetCount(t *testing.T) {
	fmt.Println(GetNetCount())
}
