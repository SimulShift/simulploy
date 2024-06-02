package cli

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
)

// Subcommand for docker compose down
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Compose down the Docker environments",
	Run: func(cmd *cobra.Command, args []string) {
		egg.NewDocker(Config.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
			Down().Compose()
	},
}
