package device

import (
	"fmt"
	"testing"
)

func TestGetInterfaceStat(t *testing.T) {
	fmt.Println(GetInterfaceStat())
}

func TestGetCounts(t *testing.T) {
	fmt.Println(GetNetCount())
}

func TestGetNetNames(t *testing.T) {
	fmt.Println(GetNetNames())
}
