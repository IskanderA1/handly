package service

import (
	"context"

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
	Create(ctx context.Context, name string) (domain.ProjectWithToken, error)
	RefreshTokens(ctx context.Context, id int64) (domain.ProjectWithToken, error)
	GetList(ctx context.Context, input ListInput) ([]domain.Project, error)
	GetById(ctx context.Context, id int64) (domain.ProjectWithToken, error)
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
	SignUp(ctx context.Context, input AdminSignUpInput, adminConfig AdminConfig) (domain.Admin, error)
	RefreshToken(ctx context.Context, refreshToken string) (domain.Session, error)
	GetList(ctx context.Context, input ListInput) ([]domain.Admin, error)
	GetByName(ctx context.Context, username string) (domain.Admin, error)
	Delete(ctx context.Context, username string) error
}

type Services struct {
	Projects Projects
	Admins   Admins
}

func NewServices(repositories *repository.Repositories, tokenManger token.Maker, config config.Config) *Services {
	projectsService := NewProjectsService(repositories.Projects, tokenManger)
	adminsService := NewAdminsService(repositories.Admins, repositories.Sessions, tokenManger, config)

	return &Services{
		Projects: projectsService,
		Admins:   adminsService,
	}
}
