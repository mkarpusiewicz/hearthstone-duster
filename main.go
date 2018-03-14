package main

import (
	"github.com/mkarpusiewicz/hearthstone-duster/cmd"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
)

func main() {
	database.Initialize()
	defer database.Close()

	cmd.Execute()
}
