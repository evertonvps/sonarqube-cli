package usecase

import (
	"errors"

	"github.com/evertonvps/sonarqube-cli/internal/entity"
)

const (
	PROJECT_ANALYSIS_TOKEN = "PROJECT_ANALYSIS_TOKEN"
)

// UserTokenInputDto model info
// @ProjectKey Project Key
// @Name Token name
// @Type Type(PROJECT_ANALYSIS_TOKEN)
type UserTokenInputDto struct {
	ProjectKey string
	Name       string
	Type       string
}

func (o *UserTokenInputDto) Validate() error {
	if len(o.ProjectKey) == 0 {
		return errors.New("--project-key is required")
	}

	if len(o.Name) == 0 {
		return errors.New("--name is required")
	}

	if len(o.Type) == 0 {
		return errors.New("--type is required")
	}

	return nil
}

type UserTokenUseCase struct {
	SonarRepository entity.UserRepository
}

func NewUserTokenUseCase(sonarRepository entity.UserRepository) *UserTokenUseCase {
	return &UserTokenUseCase{SonarRepository: sonarRepository}
}

func (u *UserTokenUseCase) CreateUserToken(input UserTokenInputDto) error {
	t := entity.NewUserToken(input.ProjectKey, input.Type, input.Name)
	err := u.SonarRepository.CreateUserToken(t)

	return err
}
