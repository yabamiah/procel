package internal

import (
	"net"
	"strconv"
	"time"

	"github.com/fatih/color"
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

func CatchPorts(host, protocol string, port int, all bool) {

	y := color.New(color.FgYellow, color.Bold)
	r := color.New(color.FgRed, color.Bold)
	g := color.New(color.FgGreen, color.Bold)

	if all {
		var ports []PortStatus

		for i := 0; i <= 1024; i++ {
			ports = append(ports, ScanPorts(host, protocol, i))
		}
		for _, port := range ports {
			y.Printf("%d: ", port.Port)

			if port.State == "Closed" {
				g.Printf("%s\n", port.State)
			} else {
				r.Printf("%s\n", port.State)
			}
		}
	} else {
		port := ScanPorts(
			host,
			protocol, 
			port)

		y.Printf("%d: ", port.Port)
		if port.State == "Closed" {
			g.Printf("%s\n", port.State)
		} else {
			r.Printf("%s\n", port.State)
		}

	}
}