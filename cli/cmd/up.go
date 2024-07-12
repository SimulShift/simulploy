package cmd

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
)

// Subcommand to bring up Docker
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up Docker environments",
	Run: func(cmd *cobra.Command, args []string) {
		detach, _ := cmd.Flags().GetBool("detach")
		profile := GetProfile()
		dockerEgg := egg.NewDocker().
			SetProfile(profile).
			SetMetaService(metaservice).
			Up()
		if detach {
			dockerEgg.Detached()
		}
		dockerEgg.Compose()
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
	upCmd.Flags().BoolP("detach", "d", false, "Run in detached mode")
}
