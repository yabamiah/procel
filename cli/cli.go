package cli

import (
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = color.RedString("Procel is an Aplicattion of Command Line Interface (CLI)")
	app.Usage = color.RedString("🔎 Search IPs, Addresses, Hosts, DNs SN and MX on internet.")
	app.Version = color.RedString("0.0.0")
	app.Author = "Yabamiah"

	flags1 :=[]cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "wikipedia.org",
		},
	}

	flags2 := []cli.Flag {
		cli.StringFlag{
			Name: "ip",
			Value: "5.255.255.50",
		},
	}

	flags3 := []cli.Flag {
		cli.StringFlag{
			Name: "protocol",
			Value: "tcp",
		},
		cli.StringFlag{
			Name: "host",
			Value: "localhost",
		},
		cli.Int64Flag{
			Name: "port",
			Value: 8080,
		},
		cli.BoolFlag{
			Name: "all",
			Hidden: true,
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ips",
			ShortName: "i",
			Usage: "Searching IPs addresses on the internet",
			Flags: flags1,
			Action: searchIPs,
		},
		{
			Name: "host",
			ShortName: "h",
			Usage: "Search hostname through ip",
			Flags: flags2,
			Action: searchHost,
		},
		{
			Name: "ns",
			ShortName: "s",
			Usage: "Search DNs serve names for a particular host",
			Flags: flags1,
			Action: searchNS,
		},
		{
			Name: "mx",
			ShortName: "m",
			Usage: "Search DNs MX for a particular host",
			Flags: flags1,
			Action: searchMX,
		},
		{
			Name: "serverall",
			ShortName: "a",
			Usage: "Serch IPs addresses and DNS SN for a particular host",
			Flags: flags1,
			Action: searchIPsNS,
		},
		{
			Name: "txt",
			ShortName: "t",
			Usage: "Search the DNS TXT",
			Flags: flags1,
			Action: searchTxt,
		},
		{
			Name: "scan",
			ShortName: "sc",
			Usage: "Scan all ports of your computer",
			Flags: flags3,
			Action: scanPort,
		},
	}

	return app
}