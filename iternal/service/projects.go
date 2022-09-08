package service

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/domain"
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

func (s *ProjectsService) Create(ctx context.Context, username string) (domain.ProjectWithToken, error) {

	if err := validator.ValidateFullName(username); err != nil {
		return domain.ProjectWithToken{}, err
	}

	param := db.CreateProjectParams{
		Name:  username,
		Token: "",
	}

	project, err := s.repository.Create(ctx, param)

	if err != nil {
		return domain.ProjectWithToken{}, err
	}

	token, _, err := s.tokenManger.CreateProjectToken(token.ProjectPayloadInput{
		ProjectId: project.ID,
		Name:      project.Name,
	})
	if err != nil {
		return domain.ProjectWithToken{}, err
	}

	updateParam := db.UpdateProjectParams{
		ID:    project.ID,
		Name:  username,
		Token: token,
	}

	res, err := s.repository.Update(ctx, updateParam)

	return domain.NewProjectWithToken(res), err
}

func (s *ProjectsService) RefreshTokens(ctx context.Context, id int64) (domain.ProjectWithToken, error) {
	res, err := s.repository.GetById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ProjectWithToken{}, fmt.Errorf("project not found")
		}
	}
	token, payload, err := s.tokenManger.CreateProjectToken(token.ProjectPayloadInput{
		ProjectId: res.ID,
		Name:      res.Name,
	})

	if err != nil {
		return domain.ProjectWithToken{}, err
	}

	param := db.UpdateProjectParams{
		ID:    payload.ProjectId,
		Name:  payload.Name,
		Token: token,
	}

	res, err = s.repository.Update(ctx, param)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ProjectWithToken{}, fmt.Errorf("project not found")
		}
	}

	return domain.NewProjectWithToken(res), err
}

func (s *ProjectsService) GetList(ctx context.Context, input ListInput) ([]domain.Project, error) {

	res, err := s.repository.GetList(ctx, db.ListProjectsParams{
		Limit:  input.Limit,
		Offset: input.Offset,
	})

	projects := make([]domain.Project, 0)

	for _, project := range res {
		projects = append(projects, domain.NewProject(project))
	}

	return projects, err
}

func (s *ProjectsService) GetById(ctx context.Context, id int64) (domain.ProjectWithToken, error) {

	res, err := s.repository.GetById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ProjectWithToken{}, fmt.Errorf("project not found")
		}
	}

	return domain.NewProjectWithToken(res), err
}

func (s *ProjectsService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("project not found")
		}
	}
	return err
}
