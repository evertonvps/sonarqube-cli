package usecase

import (
	"errors"

	"github.com/evertonvps/sonarqube-cli/internal/entity"
)

// QualityProfileInputDto model info
// @ProjectKey Project Key
// @QualityProfile Profile name
// @Language Language(java, python, ruby, etc)
type QualityProfileInputDto struct {
	ProjectKey     string `json:"projectkey"`
	QualityProfile string `json:"qualityProfile"`
	Language       string `json:"language"`
}

func (o *QualityProfileInputDto) Validate() error {
	if len(o.ProjectKey) == 0 {
		return errors.New("--project-key is required")
	}

	if len(o.QualityProfile) == 0 {
		return errors.New("--quality-profile is required")
	}

	if len(o.Language) == 0 {
		return errors.New("--language is required")
	}

	return nil
}

type QualityProfileUseCase struct {
	QualityProfileRepository entity.QualityProfileRepository
}

func NewQualityProfileUseCase(qualityProfileRepository entity.QualityProfileRepository) *QualityProfileUseCase {
	return &QualityProfileUseCase{QualityProfileRepository: qualityProfileRepository}
}

func (u *QualityProfileUseCase) Add(input QualityProfileInputDto) error {
	q := entity.NewQualityProfile(input.QualityProfile, input.Language, input.ProjectKey)
	err := u.QualityProfileRepository.AddQualityProfile(q)

	return err
}
