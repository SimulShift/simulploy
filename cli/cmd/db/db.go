package db

import (
	"github.com/simulshift/simulploy/cli/cmd"
	"github.com/spf13/cobra"
	"log"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database operations",
	Long:  `Database operations`,
}

func init() {
	log.Print("Initializing db command")
	cmd.RootCmd.AddCommand(dbCmd)
}
