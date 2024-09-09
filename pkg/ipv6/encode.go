package jjIPv6

import "net"

func Encode(input string) string {
	ip := net.ParseIP(input)

	if ip == nil {
		return ""
	}

	ipv6 := ip.To16()
	if ipv6 == nil {
		return ""
	}

	return ipv6.String()
}
