package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// Completion
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script for Zsh",
	Long: `To load completions in your current shell session: source <(simulploy completion) 
				To load completions for every new session, execute once:
				simulploy completion > "${fpath[1]}/_dockercli"`,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
