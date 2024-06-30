package cmd

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"github.com/simulshift/simulploy/egg"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "simulploy",
	Short: "Streamline docker compose commands",
	Long:  `A docker wrapper application to manage docker environments.`,
}

// Flags and their descriptions
var (
	profileFlag string
	metaservice string
)

func init() {
	// Global flags available to all subcommands
	RootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "development", "profile to use")
	if err := RootCmd.RegisterFlagCompletionFunc("profile", profileCompletionFunc); err != nil {
		fmt.Println("Error registering flag completion for --profile:", err)
		os.Exit(1)
	}
	RootCmd.PersistentFlags().StringVarP(&metaservice, "metaservice", "m", "", "choose a metaservice")
	// Register the completion function
	if err := RootCmd.RegisterFlagCompletionFunc("metaservice", metaserviceCompletion); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering completion for 'metaservice': %v\n", err)
		os.Exit(1)
	}
	carapace.Gen(RootCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func profileCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// List of profiles, potentially could be fetched from a simulConfig file, service, etc.
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
