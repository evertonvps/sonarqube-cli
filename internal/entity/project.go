package entity

import "errors"

type ProjectRepository interface {
	Create(project *Project) error
	Delete(project *Project) error
	//UpdateProjectKey(project *Project) error

}

type Project struct {
	Key       string
	Name      string
	Qualifier string
}

func NewProject(key string, name string) *Project {

	project := &Project{
		Key:  key,
		Name: name,
	}

	err := project.Validate()
	if err != nil {
		return nil
	}
	return project
}

func (p *Project) Validate() error {

	if p.Name == "" {
		return errors.New("Name is required")
	}

	if p.Key == "" {
		return errors.New("Key is required")
	}

	return nil
}
