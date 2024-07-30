package shared

import (
	"log"
	"net"
	"net/http"
)

type HttpContextAccessor struct {
	HttpContext *http.Request
}

func isIPAddressSet(address net.IP) bool {

	ipv6Loopback := net.IPv6loopback
	return address != nil && !address.Equal(ipv6Loopback)
}

func (httpContextAccessor *HttpContextAccessor) IsRequestAvailable() bool {

	if httpContextAccessor.HttpContext == nil {
		return false
	}

	if httpContextAccessor == nil {
		return false
	} // burayıı kontrol et.

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	return true
}

func IsRequestAvailable(httpContextAccessor *HttpContextAccessor) bool {
	if httpContextAccessor.HttpContext == nil {
		return false
	}

	if httpContextAccessor == nil {
		return false
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	return true
}

// GetCurrentIpAddress
func GetCurrentIpAddress(httpContextAccessor *HttpContextAccessor) string {
	if !IsRequestAvailable(httpContextAccessor) {
		return ""
	}

	remoteIp := httpContextAccessor.HttpContext.RemoteAddr
	ip, _, err := net.SplitHostPort(remoteIp)
	if err != nil {
		return ""
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ""
	}

	// IPv6 Loopback control
	if parsedIP.Equal(net.IPv6loopback) {
		return net.IPv4(127, 0, 0, 1).String()
	}

	// IPv4 convert and string response
	if ipv4 := parsedIP.To4(); ipv4 != nil {
		return ipv4.String()
	}

	return parsedIP.String()
}
