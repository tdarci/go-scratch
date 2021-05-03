package main

import (
	"fmt"
	"github.com/tdarci/prj-999/engine"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	eng := engine.NewEngine()
	var addA int
	var addB int

	app := &cli.App{
		Name:                   "test project",
		Version:                "0.001",
		Description:            "run some test commands",
		Commands: []*cli.Command{
			{
				Name: "add",
				Usage: "add two numbers",
				Flags: []cli.Flag{
						&cli.IntFlag{Name: "a", Value: 0, Destination: &addA},
						&cli.IntFlag{Name: "b", Value: 0, Destination: &addB},
				},
				Action: func(c *cli.Context) error {
					fmt.Printf("%d + %d is %d\n", addA, addB, eng.Add(addA, addB))
					return nil
				},

			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("boom boom " + ctx.Args().Get(0))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


