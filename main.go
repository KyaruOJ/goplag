package main

import (
	"fmt"
	"goplag/lexer"
	"goplag/simtest"
	"goplag/source"
	"goplag/winnowing"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "compare",
				Action: func(c *cli.Context) error {
					files := []string{
						c.Args().Get(0),
						c.Args().Get(1),
					}
					fileSrcs, err := source.Files(files)
					if err != nil {
						return err
					}

					for _, src := range fileSrcs {
						src.Tokens = lexer.Lexer(src)
						src.Fingerprints = winnowing.Winnowing(src.Tokens)
					}

					ans := simtest.Compare(fileSrcs[0], fileSrcs[1])
					fmt.Println(ans)

					return nil
				},
			},
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
