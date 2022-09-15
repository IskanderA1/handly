package token

import "fmt"

type ProjectPayloadInput struct {
	ProjectId int64
	Name      string
}

type ProjectPayload struct {
	ProjectId int64  `json:"project_id"`
	Name      string `json:"name"`
}

func NewProjectPayload(input ProjectPayloadInput) (*ProjectPayload, error) {
	payload := &ProjectPayload{
		ProjectId: input.ProjectId,
		Name:      input.Name,
	}
	return payload, nil
}

func (p *ProjectPayload) Valid() error {
	if p.ProjectId != 0 && p.Name != "" {
		return nil
	}
	return fmt.Errorf("invalid data")
}
