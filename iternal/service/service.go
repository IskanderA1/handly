package service

import (
	"context"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/repository"
	"github.com/IskanderA1/handly/pkg/config"
	"github.com/IskanderA1/handly/pkg/token"
)

type ListInput struct {
	Limit  int32
	Offset int32
}

type Projects interface {
	Create(ctx context.Context, name string) (db.Project, error)
	RefreshTokens(ctx context.Context, refreshToken string) (db.Project, error)
	GetList(ctx context.Context, input ListInput) ([]db.Project, error)
	GetById(ctx context.Context, id int64) (db.Project, error)
	Delete(ctx context.Context, id int64) error
}

type AdminSingInInput struct {
	Username string
	Password string
}

type AdminSignUpInput struct {
	Username string
	Password string
	FullName string
}

type AdminConfig struct {
	UserAgent string
	ClientIp  string
}

type Admins interface {
	SignIn(ctx context.Context, input AdminSingInInput, adminConfig AdminConfig) (domain.Session, error)
	SignUp(ctx context.Context, input AdminSignUpInput, adminConfig AdminConfig) (domain.Session, error)
	RefreshToken(ctx context.Context, refreshToken string) (domain.Session, error)
	GetList(ctx context.Context, input ListInput) ([]domain.Admin, error)
	GetByName(ctx context.Context, username string) (domain.Admin, error)
	Delete(ctx context.Context, username string) error
}

type Services struct {
	Projects Projects
}

func NewServices(repositories *repository.Repositories, tokenManger token.Maker, config config.Config) *Services {
	projectsService := NewProjectsService(repositories.Projects, tokenManger)

	return &Services{
		Projects: projectsService,
	}
}
