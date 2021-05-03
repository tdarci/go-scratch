package main

import (
	"log"

	"github.com/tdarci/prj-999/api/api"
)

const apiPort = 8181

func main() {
	s := api.NewAPI()
	err := s.Run(apiPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server shut down.")
}
