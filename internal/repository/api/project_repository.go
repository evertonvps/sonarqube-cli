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

type SonarProjectRepository struct {
	SonarCfg *entity.SonarCfg
}

func NewSonarProjectRepository(cfg *entity.SonarCfg) *SonarProjectRepository {
	return &SonarProjectRepository{SonarCfg: cfg}
}

// Create a Sonar Project
func (r *SonarProjectRepository) Create(p *entity.Project) error {

	baseUrl, err := url.Parse(fmt.Sprintf("%s/api/projects/create?", r.SonarCfg.Host))
	if err != nil {
		log.Error().Msgf("Malformed URL: %s", err.Error())
		return err
	}

	params := url.Values{}
	params.Add("project", p.Key)
	params.Add("name", p.Name)

	baseUrl.RawQuery = params.Encode()

	req, e := http.NewRequest("POST", baseUrl.String(), nil)
	if e != nil {
		log.Error().Msgf(e.Error())
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
		return err
	} else if response.StatusCode == http.StatusBadRequest {
		err = fmt.Errorf(string(b))
		return err
	} else if response.StatusCode == http.StatusOK {
		log.Info().Msgf("Created project: %s | %s",
			p.Key, p.Name)
		err := json.Unmarshal(b, &p)

		if err != nil {

			log.Error().Msgf(err.Error())

			return err
		}

	}

	return nil

}

// Delete a project.
// https://$SONAR_HOST/api/projects/
func (r *SonarProjectRepository) Delete(p *entity.Project) error {
	log.Error().Msg("Not implemented")
	return nil
}

// http://$SONAR_HOST/api/projects/updatekey
func (r *SonarProjectRepository) UpdateProjectKey() {

}
