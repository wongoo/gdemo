package network

import (
	"net"
	"strings"
)

var localhostIP = net.ParseIP("127.0.0.1")

//GetHostIPv4 get host ip address
func GetHostIPv4() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return localhostIP
	}

	var filterIP = localhostIP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		var ip net.IP
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil {
				continue
			}
			ip = ip.To4()
			if ip != nil && !ip.IsLoopback() {
				break
			}
		}

		if ip == nil {
			continue
		}
		filterIP = ip
		if !strings.Contains(i.Name, "docker") && !strings.Contains(i.Name, "bridge") {
			break
		}
	}

	return filterIP
}
