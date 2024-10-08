package cmd

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
)

// Subcommand to clean Docker images
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete Docker images for the profile",
	Run: func(cmd *cobra.Command, args []string) {
		egg.NewDocker().
			SetProfile(profileFlag).
			SetMetaService(metaservice).
			Clean().Compose()
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
