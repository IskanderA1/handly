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

type Admins interface {
	SignIn(ctx context.Context, input AdminSingInInput, adminConfig AdminConfig) (domain.Session, error)
	SignUp(ctx context.Context, input AdminSignUpInput, adminConfig AdminConfig) (domain.Admin, error)
	RefreshToken(ctx context.Context, refreshToken string) (domain.Session, error)
	GetList(ctx context.Context, input ListInput) ([]domain.Admin, error)
	GetByName(ctx context.Context, username string) (domain.Admin, error)
	Delete(ctx context.Context, username string) error
}

type Services struct {
	Projects     Projects
	Admins       Admins
	Events       Events
	ProjectsLogs ProjectsLogs
}

type Events interface {
	Create(ctx context.Context, inp CreateEventInput) (domain.Event, error)
	GetById(ctx context.Context, id int64) (domain.Event, error)
	GetListByProjectId(ctx context.Context, projectID int64) ([]domain.Event, error)
	Update(ctx context.Context, inp UpdateEventInput) (domain.Event, error)
	Delete(ctx context.Context, id int64) error
}

type ProjectsLogs interface {
	InitUser(ctx context.Context, inp UserInput) error
	SendLog(ctx context.Context, inp LogInput) error
}

type AdminLogs interface{}

type ServiceDependence struct {
	Repositories       *repository.Repositories
	AdminTokenManger   token.Maker[token.AdminPayload, token.AdminPayloadInput]
	ProjectTokenManger token.Maker[token.ProjectPayload, token.ProjectPayloadInput]
	Config             config.Config
}

func NewServices(dependence ServiceDependence) *Services {
	projectsService := NewProjectsService(dependence.Repositories.Projects, dependence.ProjectTokenManger)
	adminsService := NewAdminsService(dependence)
	eventsService := NewEventsService(dependence.Repositories.Events)
	projectsLogService := NewProjectsLogsService(
		ProjectsLogsServiceDependency{
			UserLogs:        dependence.Repositories.Users,
			LogsRepository:  dependence.Repositories.Logs,
			EventRepository: dependence.Repositories.Events,
		},
	)

	return &Services{
		Projects:     projectsService,
		Admins:       adminsService,
		Events:       eventsService,
		ProjectsLogs: projectsLogService,
	}
}
