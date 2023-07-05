package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "hubsalot",
		Commands: []*cli.Command{
			{
				Name:    "sup",
				Aliases: []string{"s"},
				Usage:   "Sir Hubsalot will list the important things",
				Action: func(c *cli.Context) error {
					println("Hello friend!")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
