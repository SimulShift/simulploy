package cmd

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/simulshift/simulploy/simulConfig"
	"github.com/spf13/cobra"
)

// Subcommand to clean Docker images
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete Docker images for the profile",
	Run: func(cmd *cobra.Command, args []string) {
		egg.NewDocker(simulConfig.Get.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
			Clean().Compose()
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
