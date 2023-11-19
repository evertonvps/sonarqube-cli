package entity

import "errors"

type QualityGateRepository interface {
	AddQualityGate(qualityGate *QualityGate) error
}

type QualityGate struct {
	Name       string
	ProjectKey string
}

func NewQualityGate(name string, projectKey string) *QualityGate {

	q := &QualityGate{
		ProjectKey: projectKey,
		Name:       name,
	}

	err := q.Validate()
	if err != nil {
		return nil
	}
	return q
}

func (q *QualityGate) Validate() error {

	if q.Name == "" {
		return errors.New("Name is required")
	}

	if q.ProjectKey == "" {
		return errors.New("Project Key is required")
	}

	return nil
}
