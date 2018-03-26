package main

import (
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/cmd"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
)

func main() {
	log.SetFlags(0) // disable time in logs

	database.Initialize()
	defer database.Close()

	cmd.Execute()
}
