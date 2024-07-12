package cmd

import (
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
		profile := GetProfile()
		dockerEgg := egg.NewDocker().
			SetProfile(profile).
			SetMetaService(metaservice).
			Down()
		if dropFlag {
			dockerEgg.Drop()
		}
		dockerEgg.Compose()
	},
}

func init() {
	downCmd.Flags().BoolVar(&dropFlag, "drop", false, "Drop all resources (e.g., remove volumes and images)")
	RootCmd.AddCommand(downCmd)
}
