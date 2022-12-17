package utils

import (
	"net"
	"time"
)

func GetOutboundIP() (net.IP, error) {
    targetAddr := "8.8.8.8:80"
    conn, err := net.DialTimeout("udp", targetAddr, 5 * time.Second)
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP, nil
}