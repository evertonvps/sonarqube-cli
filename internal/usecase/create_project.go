package usecase

import (
	"errors"

	"github.com/evertonvps/sonarqube-cli/internal/entity"
)

// CreateProjectInputDto model info
// @ProjectKey Project Key
// @ProjectName Project Name
// @QualityGate    Quality Gate name
// @QualityProfile Quality Profile name
// @Language       Language
type CreateProjectInputDto struct {
	ProjectKey     string
	ProjectName    string
	QualityGate    string
	QualityProfile string
	Language       string
}

func (o *CreateProjectInputDto) Validate() error {

	if len(o.ProjectName) == 0 {
		return errors.New("--project-name is required")
	}
	if len(o.ProjectKey) == 0 {
		return errors.New("--project-key is required")
	}

	return nil
}

// CreateProjectOutputDto model info
// @Key Project Key
// @Name Project Name
// @Qualifier Qualifier (TRK,VW,APP)
type CreateProjectOutputDto struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Qualifier string `json:"qualifier"`
}

type CreateProjectUseCase struct {
	ProjectRepository entity.ProjectRepository
}

func NewCreateProjectUseCase(projectRepository entity.ProjectRepository) *CreateProjectUseCase {
	return &CreateProjectUseCase{ProjectRepository: projectRepository}
}

func (u *CreateProjectUseCase) Execute(input CreateProjectInputDto) (*CreateProjectOutputDto, error) {

	p := entity.NewProject(input.ProjectKey, input.ProjectName)
	err := u.ProjectRepository.Create(p)

	if err != nil {
		return nil, err
	}

	return &CreateProjectOutputDto{
		Key:       p.Key,
		Name:      p.Name,
		Qualifier: p.Qualifier,
	}, nil
}
