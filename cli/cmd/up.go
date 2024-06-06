package cmd

import (
	"github.com/simulshift/simulploy/cli/config"
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
	"log"
)

// Subcommand to bring up Docker
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up Docker environments",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Running up command: ", profileFlag, metaservice)
		// print docker dir
		log.Println("Docker dir: ", config.MemoryStore.DockerDir)
		egg.NewDocker(config.MemoryStore.DockerDir).
			SetProfile(egg.Profile(profileFlag)).
			SetMetaService(egg.MetaService(metaservice)).
			Up().Compose()
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
