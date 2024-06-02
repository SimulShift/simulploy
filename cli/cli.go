package cli

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// Flags and their descriptions
var (
	profileFlag string
	metaservice string
)

func Cli() {
	// try to load the configuration
	err := Config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var rootCmd = &cobra.Command{
		Use:   "simulploy",
		Short: "Streamline docker compose commands",
		Long:  `A docker wrapper application to manage docker environments.`,
	}
	carapace.Gen(rootCmd)

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

	rootCmd.AddCommand(completionCmd)

	// Global flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&profileFlag, "profile", "development", "profile to use")
	if err := rootCmd.RegisterFlagCompletionFunc("profile", profileCompletionFunc); err != nil {
		fmt.Println("Error registering flag completion for --profile:", err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&metaservice, "metaservice", "m", "", "choose a metaservice")
	// Register the completion function
	if err := rootCmd.RegisterFlagCompletionFunc("metaservice", metaserviceCompletion); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering completion for 'metaservice': %v\n", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(dropCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
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
