package di

import (
	"fmt"
	"runtime/trace"

	"github.com/evertonvps/sonarqube-cli/internal/entity"
	"github.com/evertonvps/sonarqube-cli/internal/repository/api"
	"github.com/evertonvps/sonarqube-cli/internal/usecase"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// Create Sonar project and define quality profile/gate
func SonarProjectCmd(cfg *entity.SonarCfg) *cobra.Command {
	projectRepository := api.NewSonarProjectRepository(cfg)
	var projectUseCase = usecase.NewCreateProjectUseCase(projectRepository)
	var dto = usecase.CreateProjectInputDto{}
	var apply bool
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manager Sonar Project",
		Long:  "Manager a Sonar Project and define Quality gate and profile.",
		Example: `
#Create
sonarqube project apply --project-name=Sonar Cli --project-key=Sonar_cli_007 --quality-gate=Sonar way --quality-profile=Sonar way --language=go
#Update

#Delete
sonarqube project apply -d --project-key=Sonar_cli_007

		`,
		Args: func(cmd *cobra.Command, args []string) error {
			return dto.Validate()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			defer trace.StartRegion(cmd.Context(), "project").End()

			if apply {

				output, err := projectUseCase.Execute(dto)
				if err == nil {
					log.Debug().Msg(output.Key)
					var qualitGateDto = usecase.QualityGateInputDto{
						ProjectKey:  output.Key,
						QualityGate: dto.QualityGate,
					}

					if err := qualitGateDto.Validate(); err == nil {

						err = usecase.NewQualityGateUseCase(api.NewQualityGateRepository(cfg)).Add(qualitGateDto)
						if err != nil {
							log.Warn().Msg(err.Error())
						}
					}

					var qualitProfileDto = usecase.QualityProfileInputDto{
						ProjectKey:     output.Key,
						QualityProfile: dto.QualityProfile,
						Language:       dto.Language,
					}
					if err := qualitProfileDto.Validate(); err == nil {

						err = usecase.NewQualityProfileUseCase(api.NewQualityProfileRepository(cfg)).Add(qualitProfileDto)
						if err != nil {
							log.Warn().Msg(err.Error())
						}
					}

					var userTokenDto = usecase.UserTokenInputDto{
						ProjectKey: output.Key,
						Name:       fmt.Sprintf("Project scan: %s", output.Key),
						Type:       usecase.PROJECT_ANALYSIS_TOKEN,
					}

					err = usecase.NewUserTokenUseCase(api.NewUserRepository(cfg)).CreateUserToken(userTokenDto)
					if err != nil {
						log.Warn().Msg(err.Error())
					}

				}
				return err

			}

			return nil

		},
	}
	cmd.Flags().BoolVarP(&apply,
		"apply",
		"a",
		false, "Apply project")

	cmd.Flags().StringVarP(&dto.ProjectName,
		"project-name",
		"n",
		"", "Name of the project")
	cmd.Flags().StringVarP(&dto.ProjectKey,
		"project-key",
		"k",
		"", "Project Key")
	cmd.Flags().StringVarP(&dto.QualityGate,
		"quality-gate",
		"g",
		"", "Name of the quality gate")
	cmd.Flags().StringVarP(&dto.QualityProfile,
		"quality-profile",
		"p",
		"", "Name of the quality profile")
	cmd.Flags().StringVarP(&dto.Language,
		"language",
		"l",
		"", "Quality profile language")

	return cmd
}

// Associate a project to a quality gate.
func SonarQualityGateCmd(cfg *entity.SonarCfg) *cobra.Command {
	repository := api.NewQualityGateRepository(cfg)
	var qualityGateUseCase = usecase.NewQualityGateUseCase(repository)
	var dto = usecase.QualityGateInputDto{}
	var apply bool = false
	cmd := &cobra.Command{
		Use:   "quality-gate",
		Short: "Associate a project to a quality gate.",
		Long: `Associate a project to a quality gate.
		The 'projectId' or 'projectKey' must be provided.
		Project id as a numeric value is deprecated since 6.1. Please use the id similar to 'AU-TpxcA-iU5OvuD2FLz'.`,
		Example: `sonarqube quality-gate apply --project-key=Sonar_cli_007 --quality-gate=Sonar way`,
		Args: func(cmd *cobra.Command, args []string) error {
			return dto.Validate()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			defer trace.StartRegion(cmd.Context(), "quality-gate").End()
			if apply {
				err := qualityGateUseCase.Add(dto)
				return err
			} else {
				log.Warn().Msg("Not implemented")
			}

			return nil

		},
	}
	cmd.Flags().BoolVarP(&apply,
		"apply",
		"a",
		false, "Apply gate")
	cmd.Flags().StringVarP(&dto.ProjectKey,
		"project-key",
		"k",
		"", "Project Key")
	cmd.Flags().StringVarP(&dto.QualityGate,
		"quality-gate",
		"g",
		"", "Name of the quality gate")

	return cmd
}

// Associate a project to a quality gate.
func SonarQualityProfileCmd(cfg *entity.SonarCfg) *cobra.Command {
	repository := api.NewQualityProfileRepository(cfg)
	var qualityProfileUseCase = usecase.NewQualityProfileUseCase(repository)
	var dto = usecase.QualityProfileInputDto{}
	var apply bool

	cmd := &cobra.Command{
		Use:   "quality-profile",
		Short: "Associate a project with a quality profile.",
		Long: `Associate a project with a quality profile.
Requires one of the following permissions:
'Administer Quality Profiles'
Edit right on the specified quality profile
Administer right on the specified project`,
		Example: `sonarqube quality-gate apply --project-key=Sonar_cli_007 way --quality-profile=Sonar way --language=go`,
		Args: func(cmd *cobra.Command, args []string) error {
			return dto.Validate()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			defer trace.StartRegion(cmd.Context(), "quality-profile").End()

			if apply {
				err := qualityProfileUseCase.Add(dto)
				return err
			} else {
				log.Warn().Msg("Not implemented")
			}

			return nil

		},
	}

	cmd.Flags().BoolVarP(&apply,
		"apply",
		"a",
		false, "Apply profile")
	cmd.Flags().StringVarP(&dto.ProjectKey,
		"project-key",
		"k",
		"", "Project Key")

	cmd.Flags().StringVarP(&dto.QualityProfile,
		"quality-profile",
		"p",
		"", "Name of the quality profile")
	cmd.Flags().StringVarP(&dto.Language,
		"language",
		"l",
		"", "Quality profile language")

	return cmd
}

// Generate a user access token.
func SonarUserTokenCmd(cfg *entity.SonarCfg) *cobra.Command {
	repository := api.NewUserRepository(cfg)
	var usertTokenUseCase = usecase.NewUserTokenUseCase(repository)
	var dto = usecase.UserTokenInputDto{}

	cmd := &cobra.Command{
		Use:   "user-token",
		Short: "Generate a user access token.",
		Long: `Generate a user access token.
Please keep your tokens secret. They enable to authenticate and analyze projects.
It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user.`,
		Example: `sonarqube user-token apply --project-key=Sonar_cli_007 way --name=Analysis --type=go`,
		Args: func(cmd *cobra.Command, args []string) error {
			return dto.Validate()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			defer trace.StartRegion(cmd.Context(), "user-token").End()

			err := usertTokenUseCase.CreateUserToken(dto)
			return err

		},
	}

	cmd.Flags().StringVarP(&dto.ProjectKey,
		"project-key",
		"k",
		"", "Project Key")

	cmd.Flags().StringVarP(&dto.Name,
		"name",
		"n",
		"", "Token Description")
	cmd.Flags().StringVarP(&dto.Type,
		"type",
		"t",
		usecase.PROJECT_ANALYSIS_TOKEN, "Token Type(PROJECT_ANALYSIS_TOKEN,)")

	return cmd
}
