package cli

import (
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// Completion
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script for Zsh",
	Long: `To load completions in your current shell session: source <(simulploy completion) 
				To load completions for every new session, execute once:
				simulploy completion > "${fpath[1]}/_dockercli"`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

func profileCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// List of profiles, potentially could be fetched from a config file, service, etc.
	var completions []string
	for _, profile := range egg.ValidProfiles {
		if strings.HasPrefix(string(profile), toComplete) {
			completions = append(completions, string(profile))
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// metaserviceCompletion will provide autocomplete function for the metaservice flag
func metaserviceCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Filter based on prefix match
	var completions []string
	for _, service := range egg.ValidMetaServices {
		if strings.HasPrefix(string(service), toComplete) {
			completions = append(completions, string(service))
		}
	}
	// NoFileComp suggests that the shell should not attempt file name completion.
	return completions, cobra.ShellCompDirectiveNoFileComp
}
