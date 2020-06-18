package main

import (
	"goplag/simtest"
	"goplag/source"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "simtest",
				Action:  func(c *cli.Context) error {
					files := c.Args().Slice()
					fileSrcs, err := source.Files(files)
					if err != nil {
						return err
					}

					return simtest.Simtest(fileSrcs)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
