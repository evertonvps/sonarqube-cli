package cmd

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/evertonvps/sonarqube-cli/internal/di"
	"github.com/evertonvps/sonarqube-cli/internal/entity"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

}

func Do(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {

	rootCmd := &cobra.Command{Use: "sonarqube",
		Version: "1.0",
		Long: `sonar cli.
	   `, SilenceUsage: true,
	}
	rootCmd.PersistentFlags().BoolP("debug", "x", false, "Debugmode (default: false)")
	//TODO viper?
	var sonarCfg = entity.SonarCfg{
		Host:  "http://sonar.local",
		Token: "WNoaYWRtaW4GFkZ6Y2hhcWNyZXRl",
	}

	//rootCmd.ResetFlags()
	rootCmd.ResetCommands()
	//Project //////////////////////////////////////////////////////////////////////////////////////////////////
	var groupGitCmd = &cobra.Command{
		Use:   "project_group",
		Short: "project commands",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msg("Executing project command")
		},
	}
	// TODO Adicionar os subcomandos ao grupo e renomear ajustar o cmd do grupo
	//groupGitCmd.AddCommand(di.SonarProjectCmd(repository))
	rootCmd.AddCommand(di.SonarProjectCmd(&sonarCfg))
	rootCmd.AddCommand(groupGitCmd)

	//Quality Gate //////////////////////////////////////////////////////////////////////////////////////////////////
	rootCmd.AddCommand(di.SonarQualityGateCmd(&sonarCfg))
	//Quality profile //////////////////////////////////////////////////////////////////////////////////////////////////
	rootCmd.AddCommand(di.SonarQualityProfileCmd(&sonarCfg))
	//User //////////////////////////////////////////////////////////////////////////////////////////////////
	rootCmd.AddCommand(di.SonarUserTokenCmd(&sonarCfg))

	rootCmd.SetArgs(args)
	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		} else {
			return 1
		}
	}
	return 0
}
