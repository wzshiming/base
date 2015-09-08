package base

import (
	"net"
	"strings"
)

var LocalIP = getLocalIP()

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
