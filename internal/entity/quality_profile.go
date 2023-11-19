package entity

import "errors"

type QualityProfileRepository interface {
	AddQualityProfile(qualityProfile *QualityProfile) error
}

type QualityProfile struct {
	ProjectKey string
	Name       string
	Language   string
}

func NewQualityProfile(name string, language string, projectKey string) *QualityProfile {

	q := &QualityProfile{
		ProjectKey: projectKey,
		Name:       name,
		Language:   language,
	}

	err := q.Validate()
	if err != nil {
		return nil
	}
	return q
}

func (q *QualityProfile) Validate() error {

	if q.Name == "" {
		return errors.New("Profile Name is required")
	}

	if q.ProjectKey == "" {
		return errors.New("Project Key is required")
	}

	if q.Language == "" {
		return errors.New("Language is required")
	}

	return nil
}
