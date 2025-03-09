package utils

import (
	"fmt"
	"net"
	"os"
)

func GuessNodeIP() (net.IP, error) {
	// 方法1：尝试通过主机名解析IP
	hostname, err := os.Hostname()
	if err == nil {
		addrs, err := net.LookupIP(hostname)
		if err == nil && len(addrs) > 0 {
			for _, addr := range addrs {
				// 优先返回IPv4地址
				if ipv4 := addr.To4(); ipv4 != nil {
					return ipv4, nil
				}
			}
			// 如果没有IPv4，返回第一个地址
			return addrs[0], nil
		}
	}

	// 方法2：检查网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("can't find network interfaces: %v", err)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsLoopback() {
				continue
			}

			// 优先返回IPv4
			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4, nil
			}
		}
	}

	return nil, fmt.Errorf("guess node ip failed")
}
