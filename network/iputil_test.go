package network

import (
	"fmt"
	"testing"
)

func TestIpUtil(t *testing.T) {
	ip := GetHostIPv4()
	fmt.Println(ip)
}
