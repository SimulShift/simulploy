package cmd

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"github.com/simulshift/simulploy/egg"
	"github.com/simulshift/simulploy/simulConfig"
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
	devMode     bool
	prodMode    bool
)

func init() {
	simulConfig.Get.Hydrate()

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
	RootCmd.PersistentFlags().BoolVarP(&devMode, "dev", "D", false, "development build")
	RootCmd.PersistentFlags().BoolVarP(&prodMode, "prod", "P", false, "production build")
	// check to make sure that only one mode is selected
	RootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if devMode && prodMode {
			log.Fatal("Cannot specify both development and production mode")
		}
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
	for _, profile := range egg.Profiles {
		if strings.HasPrefix(profile, toComplete) {
			completions = append(completions, string(profile))
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// metaserviceCompletion will provide autocomplete function for the metaservice flag
func metaserviceCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var completions []string
	for _, service := range simulConfig.Get.Metaservices {
		if strings.HasPrefix(service, toComplete) {
			completions = append(completions, service)
		}
	}
	// NoFileComp suggests that the shell should not attempt file name completion.
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetProfile() string {
	if devMode {
		return "development"
	}
	if prodMode {
		return "production"
	}
	return profileFlag
}
