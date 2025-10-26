package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"starterA/internal/app"
	"starterA/internal/database/repo"
)

type Service struct {
	Ctx    context.Context
	App    *app.Application
	Logger *slog.Logger
}

func New(ctx context.Context, app *app.Application, logger *slog.Logger) *Service {

	if logger == nil {
		if app.Logger == nil {
			logger = slog.Default()
		} else {
			logger = app.Logger
		}
	}

	return &Service{
		Ctx:    ctx,
		App:    app,
		Logger: logger,
	}
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (s *Service) CreateUser(ctx context.Context, input *CreateUserInput) (*repo.User, error) {
	s.Logger.Info("Creating user", "username", input.Username, "email", input.Email)

	user, err := s.App.DB.CreateUser(ctx, repo.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
	})
	if err != nil {
		s.Logger.Error("Failed to create user", "error", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (s *Service) GetUsers(ctx context.Context, limit, offset int64) ([]repo.User, error) {
	s.Logger.Info("Fetching users", "limit", limit, "offset", offset)

	users, err := s.App.DB.ListUsers(ctx, repo.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		s.Logger.Error("Failed to fetch users", "error", err)
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	return users, nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*repo.User, error) {
	s.Logger.Info("Fetching user", "id", id)

	user, err := s.App.DB.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		s.Logger.Error("Failed to fetch user", "error", err)
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (s *Service) UpdateUser(ctx context.Context, id int64, input *UpdateUserInput) (*repo.User, error) {
	s.Logger.Info("Updating user", "id", id)

	user, err := s.App.DB.UpdateUser(ctx, repo.UpdateUserParams{
		ID:       id,
		Username: input.Username,
		Email:    input.Email,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		s.Logger.Error("Failed to update user", "error", err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	s.Logger.Info("Deleting user", "id", id)

	err := s.App.DB.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user not found")
		}
		s.Logger.Error("Failed to delete user", "error", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
