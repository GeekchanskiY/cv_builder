package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, router *httprouter.Router) *http.Server {

	srv := &http.Server{Addr: fmt.Sprintf(":%s", os.Getenv("server_port")), Handler: router}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
