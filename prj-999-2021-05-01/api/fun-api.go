package main

import (
	"log"

	"github.com/tdarci/prj-999/api/api"
	"github.com/tdarci/prj-999/config"
)

const apiPort = 8181

func main() {
	cfg, err := config.NewLocal()
	if err != nil {
		log.Fatalf("Error creating config: %s", err)
	}
	s := api.NewAPI(cfg)
	err = s.Run(apiPort)
	if err != nil {
		cfg.Logger().Fatal(err)
	}
	cfg.Logger().Println("Server shut down.")
}
