package db

import (
	"github.com/simulshift/simulploy/simulSsh"
	"github.com/spf13/cobra"
	"log"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the database",
	Long:  `Push the database to the remote server`,
	Run: func(cmd *cobra.Command, args []string) {
		sssh := simulSsh.New()
		commandStr := "cd ~/simulchatbot && npx drizzle-kit push"
		res, err := sssh.Exec(commandStr)
		if err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
		log.Printf("Result: %s", res)
		sssh.Close()
	},
}

func init() {
	dbCmd.AddCommand(pushCmd)
}
