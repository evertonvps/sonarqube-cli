package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/evertonvps/sonarqube-cli/internal/entity"

	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	SonarCfg *entity.SonarCfg
}

func NewUserRepository(cfg *entity.SonarCfg) *UserRepository {
	return &UserRepository{SonarCfg: cfg}
}

// https://$SONAR_HOST/api/user_tokens/generate
func (r *UserRepository) CreateUserToken(u *entity.UserToken) error {

	baseUrl, err := url.Parse(fmt.Sprintf("%s/api/user_tokens/generate?", r.SonarCfg.Host))
	if err != nil {
		log.Error().Msgf("Malformed URL: %s", err.Error())
		return err
	}

	params := url.Values{}
	params.Add("name", fmt.Sprintf("Project scan on %s", u.Name))
	params.Add("projectKey", u.ProjectKey)
	params.Add("type", u.Type)

	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", baseUrl.String(), nil)

	if err != nil {
		log.Error().Msgf(err.Error())
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", r.SonarCfg.Token))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error().Msgf(err.Error())
		return err
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)

	if err != nil {
		log.Error().Msgf(err.Error())
	} else if response.StatusCode == http.StatusBadRequest {
		err = fmt.Errorf(string(b))
		log.Error().Msgf(err.Error())
	} else if response.StatusCode == http.StatusOK {

		err := json.Unmarshal(b, &u)

		if err != nil {

			log.Error().Msgf(err.Error())

			return err
		}
		log.Info().Msgf("Token %s generated: %s", u.Type, u.Token)
	}

	return nil

}
