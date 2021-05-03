package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tdarci/prj-999/config"
	"github.com/tdarci/prj-999/engine"
)

func main() {

	cfg, err := config.NewLocal()
	if err != nil {
		log.Fatalf("Error creating config: %s", err)
	}
	eng := engine.NewEngine(cfg)
	var addA int
	var addB int

	app := &cli.App{
		Name:        "test project",
		Version:     "0.001",
		Description: "run some test commands",
		Commands: []*cli.Command{
			{
				Name:  "add",
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
	err = app.Run(os.Args)
	if err != nil {
		cfg.Logger().Fatal(err)
	}
}
