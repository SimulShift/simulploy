package cli

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
)

// Drop command to clean postgres database
var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database",
	Run: func(cmd *cobra.Command, args []string) {
		egg.NewDocker(Config.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
			Drop().Compose()
	},
}
