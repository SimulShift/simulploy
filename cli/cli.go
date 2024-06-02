package cli

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// Flags and their descriptions
var (
	rootCmd     *cobra.Command
	profileFlag string
	metaservice string
)

func Cli() {
	// try to load the configuration
	err := Config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	rootCmd = &cobra.Command{
		Use:   "simulploy",
		Short: "Streamline docker compose commands",
		Long:  `A docker wrapper application to manage docker environments.`,
	}

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

	// completion commands
	rootCmd.AddCommand(completionCmd)
	carapace.Gen(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
