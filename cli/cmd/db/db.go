package db

import (
	"github.com/simulshift/simulploy/cli/cmd"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database operations",
	Long:  `Database operations`,
}

func init() {
	cmd.RootCmd.AddCommand(dbCmd)
}
