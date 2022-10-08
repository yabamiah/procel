package cli

import (
	"github.com/urfave/cli"
	"github.com/fatih/color"
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
	}

	return app
}