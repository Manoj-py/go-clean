package main

import (
	"github.com/Manoj-py/backend-architecture/common/sqlConnectors"
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/handlers"
	"log/slog"
	"net/http"

	httprouters "github.com/Manoj-py/backend-architecture/common/httprouter"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/routers"
)

type Config struct {
	routes http.Handler
}

func main() {
	//initialise the config

	//initialise the sql-db
	conn := sqlConnectors.CreateNewPostgresConnection("postgresql://postgres:password@localhost:5432/taskmanager?sslmode=disable")
	defer conn.Close()

	//initialise for sqlc
	store := db.NewStore(conn)

	//initialise the redis-db

	//initialise any other if required

	// initialise the token

	//initialise the event handler if required

	// initialised the router

	// initialised the handler
	handler := handlers.NewHandler(store)

	mux := &httprouters.ChiRouter{}
	routes := routers.Routes(mux, handler)

	if err := listenAndServe(routes); err != nil {
		slog.Error(err.Error())
		return
	}
}

func listenAndServe(routes http.Handler) error {
	server := http.Server{
		Addr:    ":8083",
		Handler: routes,
	}

	slog.Info("Started the ther server on port:" + "8083")
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Error in starting the server")
		return err
	}
	return nil

}
