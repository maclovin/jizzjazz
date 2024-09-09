package jjIPv6

import "net"

func Decode(input string) string {
	ip := net.ParseIP(input)

	if ip == nil {
		return ""
	}

	if ipv4 := ip.To4(); ipv4 != nil {
		return ipv4.String()
	}

	return ""
}
