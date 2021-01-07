package onet

/**
 * @DateTime   : 2020/12/30
 * @Author     : xumamba
 * @Description:
 **/

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	host, ip := GetHost()
	fmt.Println(host, ip)
}
