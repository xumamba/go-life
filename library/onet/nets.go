package onet

/**
 * @DateTime   : 2020/12/30
 * @Author     : xumamba
 * @Description:
 **/

import (
	"net"
	"os"
)

func GetHost() (hostName, IP string) {
	hostName, _ = os.Hostname()
	ips, _ := net.LookupIP(hostName)
	for _, ip := range ips{
		if ipv4 := ip.To4(); ipv4 != nil{
			return hostName, ipv4.String()
		}
	}
	return hostName, ""
}
