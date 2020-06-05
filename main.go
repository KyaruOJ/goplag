package main

import (
	"fmt"
	"goplag/simtest"
	"goplag/source"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var raw bool
	var comparing string
	var compared string
	var orig bool

	app := &cli.App{
		Name:  "goplag",
		Usage: "A go program for code plagiarism detection.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "raw",
				Value:       false,
				Required:    false,
				Destination: &raw,
			},
			&cli.BoolFlag{
				Name:        "orig",
				Value:       false,
				Required:    false,
				Destination: &orig,
			},
			&cli.StringFlag{
				Name:        "comparing",
				Aliases:     []string{"a"},
				Required:    true,
				Destination: &comparing,
			},
			&cli.StringFlag{
				Name:        "compared",
				Aliases:     []string{"b"},
				Required:    true,
				Destination: &compared,
			},
		},
		Action: func(c *cli.Context) error {
			var ans string
			var err error

			aSrc, err := source.Open(comparing)
			if err != nil {
				return err
			}
			bSrc, err := source.Open(compared)
			if err != nil {
				return err
			}

			if orig {
				ans, err = simtest.Origtest(aSrc, bSrc)
			} else {
				ans, err = simtest.Simtest(aSrc, bSrc, raw)
			}
			if err != nil {
				return err
			}

			fmt.Println(ans)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
