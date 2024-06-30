package cmd

import (
	"fmt"
	"github.com/simulshift/simulploy/simulConfig"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "simulConfig",
	Short: "Configure the CLI",
	Long:  `Configure the CLI with the necessary settings`,
	Run: func(cmd *cobra.Command, args []string) {
		// do help here
		fmt.Println("Usage: simulploy simulConfig [get|set|save]")
	},
}

var saveConfigCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the configuration",
	Long:  `Save the configuration to a yaml file, .simulploy.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		simulConfig.Get.Save()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(saveConfigCmd)
}
