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
	var base string
	var comparing string
	var compared string

	app := &cli.App{
		Name:  "goplag",
		Usage: "A go program for code plagiarism detection.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "base",
				Value:       "",
				Required:    false,
				Destination: &base,
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
			var baseSrc *source.Source = nil
			var err error
			if base != "" {
				baseSrc, err = source.Open(base)
				if err != nil {
					return err
				}
			}

			aSrc, err := source.Open(comparing)
			if err != nil {
				return err
			}
			bSrc, err := source.Open(compared)
			if err != nil {
				return err
			}

			ans, err := simtest.Simtest(baseSrc, aSrc, bSrc)
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
