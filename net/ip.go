package net

import (
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"
)

/*
 * Ipv4 2 Int
 */
func Ipv4ToInt(clientIp string) int64 {
	var intIP int64 = 0
	ipStrArr := strings.Split(clientIp, ".")
	ipStrArrLen := len(ipStrArr)
	for k, v := range ipStrArr {
		intV, _ := strconv.Atoi(v)

		intIP += int64(int(math.Pow(256, float64(ipStrArrLen-1-k))) * intV)
	}
	return intIP
}

/*
 * Int 2 Ipv4
 */
func IntToIpv4(ipInt int64) string {
	var ipv4Address string = ""
	var segment string = ""

	for ipInt > 0 {
		segment = strconv.Itoa(int(ipInt % 256))
		ipv4Address = segment + "." + ipv4Address
		intSegment, _ := strconv.Atoi(segment)
		ipInt -= int64(intSegment)
		ipInt /= 256
	}

	ipv4Address = strings.TrimRight(ipv4Address, ".")

	return ipv4Address
}

/*
 * GetClientIP
 */
func GetClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	remoteAddr := strings.TrimSpace(r.RemoteAddr)
	if ip, _, err := net.SplitHostPort(remoteAddr); err == nil {
		return ip
	}

	return ""
}
