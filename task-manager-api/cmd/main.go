package main

import (
	"github.com/Manoj-py/backend-architecture/common/environment"
	"github.com/Manoj-py/backend-architecture/common/sqlConnectors"
	"github.com/Manoj-py/backend-architecture/task-manager-api/cmd/config"
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/handlers"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/service"
	"log/slog"
	"net/http"

	httprouters "github.com/Manoj-py/backend-architecture/common/httprouter"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/routes"
)

type Server struct {
	routes http.Handler
	config config.Config
}

func main() {
	//initialise the config
	environmentI := environment.NewLoadConfig()
	env := config.MustLoadConfig(environmentI)

	//initialise the sql-db
	conn := sqlConnectors.CreateNewPostgresConnection(env.POSTGRES_URL)
	defer conn.Close()

	//initialise for sqlc
	repo := db.NewStore(conn)

	//initialise the service layer and give repo to it
	svc := service.NewService(repo)

	//initialise the redis-db

	//initialise any other if required

	// initialise the token

	//initialise the message-broker handler if required

	// initialised the handler
	handler := handlers.NewHandler(repo, svc)

	// initialised the router
	router := httprouters.NewRouter()
	mux := routes.Routes(router, handler)

	if err := listenAndServe(mux, env); err != nil {
		slog.Error(err.Error())
		return
	}
}
