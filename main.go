package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	app := &cli.App{
		Name:                   "kp",
		Usage:                  "Kill the process running on given port",
		Action: func(context *cli.Context) error {
			port := context.Args().Get(0)
			KillPort(port)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}

}