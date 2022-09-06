package service

import (
	"context"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/repository"
	"github.com/IskanderA1/handly/pkg/token"
	"github.com/IskanderA1/handly/pkg/validator"
)

type ProjectsService struct {
	repository  repository.Projects
	tokenManger token.Maker
}

func NewProjectsService(repository repository.Projects, tokenManger token.Maker) *ProjectsService {
	return &ProjectsService{
		repository:  repository,
		tokenManger: tokenManger,
	}
}

func (s *ProjectsService) Create(ctx context.Context, username string) (db.Project, error) {

	if err := validator.ValidateFullName(username); err != nil {
		return db.Project{}, err
	}

	param := db.CreateProjectParams{
		Name:  username,
		Token: "",
	}

	project, err := s.repository.Create(ctx, param)

	if err != nil {
		return db.Project{}, err
	}

	token, _, err := s.tokenManger.CreateProjectToken(token.ProjectPayloadInput{
		ProjectId: project.ID,
		Name:      project.Name,
	})
	if err != nil {
		return db.Project{}, err
	}

	updateParam := db.UpdateProjectParams{
		ID:    project.ID,
		Name:  username,
		Token: token,
	}

	return s.repository.Update(ctx, updateParam)
}

func (s *ProjectsService) RefreshTokens(ctx context.Context, refreshToken string) (db.Project, error) {
	oldPayload, err := s.tokenManger.VerifyProjectToken(refreshToken)
	if err != nil {
		return db.Project{}, err
	}

	token, payload, err := s.tokenManger.CreateProjectToken(token.ProjectPayloadInput{
		ProjectId: oldPayload.ProjectId,
		Name:      oldPayload.Name,
	})

	if err != nil {
		return db.Project{}, err
	}

	param := db.UpdateProjectParams{
		ID:    payload.ProjectId,
		Name:  payload.Name,
		Token: token,
	}

	return s.repository.Update(ctx, param)
}

func (s *ProjectsService) GetList(ctx context.Context, input ListInput) ([]db.Project, error) {
	return s.repository.GetList(ctx, db.ListProjectsParams{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
}

func (s *ProjectsService) GetById(ctx context.Context, id int64) (db.Project, error) {
	return s.repository.GetById(ctx, id)
}

func (s *ProjectsService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
