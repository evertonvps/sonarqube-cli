package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/evertonvps/sonarqube-cli/internal/entity"

	"github.com/rs/zerolog/log"
)

type QualityProfileRepository struct {
	SonarCfg *entity.SonarCfg
}

func NewQualityProfileRepository(cfg *entity.SonarCfg) *QualityProfileRepository {
	return &QualityProfileRepository{SonarCfg: cfg}
}

// https://$SONAR_HOST/api/qualityprofiles/add_project
func (r *QualityProfileRepository) AddQualityProfile(q *entity.QualityProfile) error {

	baseUrl, err := url.Parse(fmt.Sprintf("%s/api/qualityprofiles/add_project?", r.SonarCfg.Host))
	if err != nil {
		log.Error().Msgf("Malformed URL: %s", err.Error())
		return err
	}

	params := url.Values{}
	params.Add("project", q.ProjectKey)
	params.Add("language", q.Language)
	params.Add("qualityProfile", q.Name)

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
	} else if response.StatusCode == http.StatusNoContent {
		log.Info().Msgf("Defining the project's(%s) quality profile %s", q.ProjectKey, q.Name)

	}
	return err
}
