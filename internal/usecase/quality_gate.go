package usecase

import (
	"errors"

	"github.com/evertonvps/sonarqube-cli/internal/entity"
)

// QualityGateInputDto model info
// @ProjectKey Project Key
// @ProjectName Project Name
type QualityGateInputDto struct {
	ProjectKey  string `json:"projectkey"`
	QualityGate string `json:"qualitygate"`
}

func (o *QualityGateInputDto) Validate() error {
	if len(o.ProjectKey) == 0 {
		return errors.New("--project-key is required")
	}

	if len(o.QualityGate) == 0 {
		return errors.New("--quality-gate is required")
	}

	return nil
}

type QualityGateUseCase struct {
	QualityGateRepository entity.QualityGateRepository
}

func NewQualityGateUseCase(qualityGateRepository entity.QualityGateRepository) *QualityGateUseCase {
	return &QualityGateUseCase{QualityGateRepository: qualityGateRepository}
}

func (u *QualityGateUseCase) Add(input QualityGateInputDto) error {
	q := entity.NewQualityGate(input.QualityGate, input.ProjectKey)
	err := u.QualityGateRepository.AddQualityGate(q)

	if err != nil {
		return err
	}

	return nil
}
