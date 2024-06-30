package cmd

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/simulshift/simulploy/simulConfig"
	"github.com/spf13/cobra"
	"log"
)

// Subcommand to bring up Docker
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up Docker environments",
	Run: func(cmd *cobra.Command, args []string) {
		detach, _ := cmd.Flags().GetBool("detach")
		log.Println("Running up command: ", profileFlag, metaservice)
		// print docker dir
		log.Println("Docker dir: ", simulConfig.Get.DockerDir)
		dockerEgg := egg.NewDocker(simulConfig.Get.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
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
