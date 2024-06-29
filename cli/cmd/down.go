package cmd

import (
	"github.com/simulshift/simulploy/cli/config"
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
)

// Declare the flag variable
var dropFlag bool

// Subcommand for docker compose down
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Compose down the Docker environments",
	Run: func(cmd *cobra.Command, args []string) {
		dockerEgg := egg.NewDocker(config.MemoryStore.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
			Down()
		if dropFlag {
			dockerEgg.Drop()
		}
		dockerEgg.Compose()
	},
}

func init() {
	downCmd.Flags().BoolVar(&dropFlag, "drop", false, "Drop all resources (e.g., remove volumes and images)")
	rootCmd.AddCommand(downCmd)
}
