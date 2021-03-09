// Package util...
//
// Description : util...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 5:56 下午
package util

import "net"

// GetHostIP 获取本机IP地址
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:58 下午 2021/3/9
func GetHostIP() string {
	hostIP := "127.0.0.1"
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				hostIP = ipnet.IP.String()
				break
			}
		}
	}
	return hostIP
}
