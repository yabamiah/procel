package cli

import (
	"github.com/urfave/cli"
)

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicattion of Command Line Interface (CLI)"
	app.Usage = "Search IPs, Addresses, Hosts, DNs SN and MX on internet." 

	flags1 :=[]cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "amazon.com.br",
		},
	}

	flags2 := []cli.Flag {
		cli.StringFlag{
			Name: "ip",
			Value: "5.255.255.50",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Searching IP addresses on the internet",
			Flags: flags1,
			Action: searchIPs,
		},
		{
			Name: "host",
			Usage: "Search hostname through ip",
			Flags: flags2,
			Action: searchHost,
		},
		{
			Name: "sn",
			Usage: "Search DNs serve names",
			Flags: flags1,
			Action: searchNS,
		},
		{
			Name: "mx",
			Usage: "Search DNs MX for the given domain nam",
			Flags: flags1,
			Action: searchMX,
		},
	}

	return app
}