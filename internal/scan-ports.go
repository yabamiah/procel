package internal

import (
	"net"
	"strconv"
	"time"
)

type PortStatus struct {
	Port  int
	State string
}

func ScanPorts(host, protocol string, port int) PortStatus {
	status := PortStatus{Port: port}
	address := host + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		status.State = "Closed"
		return status
	}
	defer conn.Close()
	
	status.State = "Open"
	return status
}