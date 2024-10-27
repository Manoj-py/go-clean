package service

import (
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/models"
	"golang.org/x/net/context"
)

type ServiceI interface {
	CreateUser(ctx context.Context, user models.UserRequestModel) (db.User, error)
}

type Service struct {
	repo db.Store
}

func NewService(repo db.Store) ServiceI {
	return &Service{
		repo: repo,
	}
}
