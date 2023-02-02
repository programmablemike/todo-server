package cmd

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	todov1 "github.com/programmablemike/todo-api/autogen/go/todo/v1"
	"github.com/programmablemike/todo-api/autogen/go/todo/v1/todov1connect"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TodoServer struct{}

func (ts *TodoServer) CreateTask(
	ctx context.Context,
	req *connect.Request[todov1.CreateTaskRequest],
) (*connect.Response[todov1.CreateTaskResponse], error) {
	return nil, nil
}

func (ts *TodoServer) DeleteTask(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTaskRequest],
) (*connect.Response[todov1.DeleteTaskResponse], error) {
	return nil, nil
}

func (ts *TodoServer) ListTasks(
	ctx context.Context,
	req *connect.Request[todov1.ListTasksRequest],
) (*connect.Response[todov1.ListTasksResponse], error) {
	return nil, nil
}

func (ts *TodoServer) MarkTask(
	ctx context.Context,
	req *connect.Request[todov1.MarkTaskRequest],
) (*connect.Response[todov1.MarkTaskResponse], error) {
	return nil, nil
}

func NewServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "launches the Todo API server",
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("launching the server")
			mux := http.NewServeMux()
			// The generated constructors return a path and a plain net/http
			// handler.
			mux.Handle(todov1connect.NewTodoServiceHandler(&TodoServer{}))
			err := http.ListenAndServe(
				"localhost:8080",
				// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
				// avoid x/net/http2 by using http.ListenAndServeTLS.
				h2c.NewHandler(mux, &http2.Server{}),
			)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
