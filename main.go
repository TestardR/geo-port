package main

import (
	"github.com/TestardR/geo-port/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	consoleOutput := log.New(os.Stdout, "", log.LstdFlags)

	app := &cli.App{
		Name:  "geo-port",
		Usage: "CLI application for managing GP service",
		Commands: []*cli.Command{
			{
				Name: "Add or Update Ports",
				Usage: "To add or update ports, you should provide a file in json format with ports data. " +
					"For example: aup -f my_port_info.json.",
				Aliases: []string{"aup"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: cmd.AddOrUpdateFilePathArgument, Aliases: []string{"f"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					return cmd.RunAsCLIHandler(c, consoleOutput)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		consoleOutput.Fatal(err)
	}
}
