package service

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/repository"
)

type UserInput struct {
	ProjectAccountID string
	Name             string
	Uuid             string
}

type LogInput struct {
	ProjectId int64
	EventName string
	UserUUID  string
	Data      sql.NullString
}

type ProjectsLogsServiceDependency struct {
	UserLogs        repository.Users
	LogsRepository  repository.Logs
	EventRepository repository.Events
}

type ProjectsLogsService struct {
	userLogs        repository.Users
	logsRepository  repository.Logs
	eventRepository repository.Events
}

func NewProjectsLogsService(d ProjectsLogsServiceDependency) *ProjectsLogsService {
	return &ProjectsLogsService{
		userLogs:        d.UserLogs,
		logsRepository:  d.LogsRepository,
		eventRepository: d.EventRepository,
	}
}

func (s *ProjectsLogsService) InitUser(ctx context.Context, inp UserInput) error {
	if inp.ProjectAccountID == "" && inp.Uuid == "" {
		return fmt.Errorf("accountId and uuid is empty")
	}
	userOrNullByAccountId, err := s.userLogs.GetUserByProjectAccountId(ctx, inp.ProjectAccountID)
	if err != nil {
		return err
	}
	userOrNullByUUID, err := s.userLogs.GetUserByUUID(ctx, inp.Uuid)
	if err != nil {
		return err
	}
	if !userOrNullByAccountId.Valid && !userOrNullByUUID.Valid {
		fmt.Println(inp)
		_, err := s.userLogs.Create(ctx, db.CreateUserParams{
			ProjectAccountID: sql.NullString{
				String: inp.ProjectAccountID,
				Valid:  inp.ProjectAccountID != "",
			},
			Uuid: sql.NullString{
				String: inp.Uuid,
				Valid:  inp.Uuid != "",
			},
			Name: sql.NullString{
				String: inp.Name,
				Valid:  inp.Name != "",
			},
		})
		if err != nil {
			return err
		}
		return nil
	}
	if userOrNullByAccountId.Valid {
		err = s.updateUser(ctx, userOrNullByAccountId.User, inp)
		if err != nil {
			return err
		}
	} else if userOrNullByUUID.Valid {
		err = s.updateUser(ctx, userOrNullByUUID.User, inp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ProjectsLogsService) updateUser(ctx context.Context, user db.User, inp UserInput) error {
	projectAccountID := user.ProjectAccountID.String
	uuid := user.Uuid.String
	name := user.Name.String
	if inp.ProjectAccountID != "" {
		projectAccountID = inp.ProjectAccountID
	}
	if inp.Uuid != "" {
		uuid = inp.Uuid
	}
	if inp.Name != "" {
		name = inp.Name
	}
	userOrNull, err := s.userLogs.Update(ctx, db.UpdateUserParams{
		ProjectAccountID: sql.NullString{
			String: projectAccountID,
			Valid:  projectAccountID != "",
		},
		Uuid: sql.NullString{
			String: uuid,
			Valid:  uuid != "",
		},
		Name: sql.NullString{
			String: name,
			Valid:  name != "",
		},
	})
	if !userOrNull.Name.Valid {
		return fmt.Errorf("User not found")
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectsLogsService) SendLog(ctx context.Context, inp LogInput) error {

	event, err := s.eventRepository.GetByName(ctx, inp.EventName)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("event not found")
		}
		return err
	}

	userOrNullByUUID, err := s.userLogs.GetUserByUUID(ctx, inp.UserUUID)
	if err != nil {
		return err
	}
	if !userOrNullByUUID.Valid {
		return fmt.Errorf("user not initizilized")
	}

	_, err = s.logsRepository.Create(ctx, db.CreateLogParams{
		ProjectID: inp.ProjectId,
		EventID:   event.ID,
		UserID:    userOrNullByUUID.User.ID,
		Data:      inp.Data,
	})

	return err
}
