package entity

import "errors"

type UserRepository interface {
	CreateUserToken(userToken *UserToken) error
}

type UserToken struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	CreatedAt  string `json:"createdAt"`
	Type       string `json:"type"`
	ProjectKey string `json:"projectKey"`
}

func NewUserToken(projectKey string, tokenType string, name string) *UserToken {

	token := &UserToken{
		ProjectKey: projectKey,
		Name:       name,
		Type:       tokenType,
	}

	err := token.Validate()
	if err != nil {
		return nil
	}
	return token
}

func (p *UserToken) Validate() error {

	if p.Name == "" {
		return errors.New("Name is required")
	}

	if p.ProjectKey == "" {
		return errors.New("Project Key is required")
	}

	if p.Type == "" {
		return errors.New("Type is required")
	}

	return nil
}
