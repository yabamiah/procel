package cli

import (
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Init() *cli.App {
	r := color.New(color.FgRed, color.Bold)

	app := cli.NewApp()
	app.Name = r.Sprint("Procel is an Aplicattion of Command Line Interface (CLI)")
	app.Usage = r.Sprint("🔎 Search IPs, Addresses, Hosts, DNs SN and MX on internet.")
	app.Version = r.Sprint("0.0.0")
	app.Author = "Yabamiah"
	app.EnableBashCompletion = true

	// Flags 1 is for commands need a host value
	flags1 :=[]cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "wikipedia.org",
		},
	}

	// Flags 2 is for commands need a ip value
	flags2 := []cli.Flag {
		cli.StringFlag{
			Name: "ip",
			Value: "5.255.255.50",
		},
	}

	// Flags 3 is for the scan command
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

	flags4 := []cli.Flag {
		cli.StringFlag{
			Name: "file",
			Value: "file.txt",
		},
		cli.StringFlag{
			Name: "size",
			Value: "100",
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
			ShortName: "n",
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
			ShortName: "si",
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
			ShortName: "s",
			Usage: "Scan all ports of your computer",
			Flags: flags3,
			Action: scanPort,
		},
		{
			Name: "archive",
			ShortName: "a",
			Usage: "Insert characters randomly in a archive",
			Flags: flags4,
			Action: archive,
		},
	}

	return app
}