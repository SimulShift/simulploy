package cli

import (
	"github.com/simulshift/simulploy/cli/cmd"
	"github.com/simulshift/simulploy/cli/config"
	"log"
)

func Cli() {
	// try to load the configuration
	err := config.MemoryStore.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	cmd.Execute()
}
