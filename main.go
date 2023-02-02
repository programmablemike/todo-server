package main

import (
	"os"

	"github.com/programmablemike/todo-server/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	app := cmd.NewTodoApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
