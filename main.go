package main

import (
	"log"
	"os"

	"github.com/yabamiah/procel/cli"
)

func main() {
	applicattion := cli.Init()
	
	err := applicattion.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}