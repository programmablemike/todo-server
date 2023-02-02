package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewTodoApp() *cli.App {
	return &cli.App{
		Name:  "todo",
		Usage: "backend server for the Todo API",
		Commands: []*cli.Command{
			NewServerCommand(),
		},
	}
}
