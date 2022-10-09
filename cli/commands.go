package cli

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	catch(err)

	color.Red("IPs:")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchHost(c *cli.Context) {
	ip := c.String("ip")

	hosts, err := net.LookupAddr(ip)
	catch(err)

	color.Red("Host:")
	for _, host := range hosts {
		fmt.Println(host)
	}
}

func searchNS(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	catch(err)

	color.Red("DNS NS:")
	for _, server := range servers {
		fmt.Println(server)
	}

}

func searchMX(c *cli.Context) {
	host := c.String("host")

	mxs, err := net.LookupMX(host)
	catch(err)

	color.Red("DNS MS:")
	for _, mx := range mxs {
		fmt.Println(mx.Host, mx.Pref)
	}
}

func searchIPsNS(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	catch(err)

	color.Red("IPs:")
	for _, ip := range ips {
		fmt.Println(ip)
	}

	servers, err := net.LookupNS(host)
	catch(err)

	color.Red("DNS NS:")
	for _, server := range servers {
		fmt.Println(server)
	}
}

func searchTxt(c *cli.Context) {
	host := c.String("host")

	txts, err := net.LookupTXT(host)
	catch(err)

	color.Red("CNAME:")
	for _, txt := range txts {
		fmt.Println(txt)
	} 
}