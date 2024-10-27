package main

import (
	"fmt"
	"github.com/Manoj-py/backend-architecture/task-manager-api/cmd/config"
	"log/slog"
	"net/http"
)

func listenAndServe(routes http.Handler, env *config.Config) error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: routes,
	}
	fmt.Println(env.PORT)
	slog.Info("Started the the server on port: " + env.PORT)
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Error in starting the server")
		return err
	}
	return nil

}
