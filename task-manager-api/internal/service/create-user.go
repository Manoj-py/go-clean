package service

import (
	"context"
	"github.com/Manoj-py/backend-architecture/common/util"
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx context.Context, user models.UserRequestModel) (db.User, error) {
	userId, err := uuid.NewUUID()
	if err != nil {
		return db.User{}, err
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return db.User{}, err
	}
	res, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Userid:       userId,
		Name:         user.Name,
		Email:        user.Email,
		Password:     hashedPassword,
		Role:         db.NullUserRole{},
		Profilephoto: "",
	})
	if err != nil {
		return db.User{}, util.ErrInternal
	}

	return res, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", util.ErrHasing
	}
	return string(hashedPassword), nil
}
