package main

import (
	"github.com/simulshift/simulploy/cli"
	"github.com/simulshift/simulploy/simulConfig"
	"log"
)

func main() {
	// try to load the configuration
	err := simulConfig.Get.Hydrate()
	if err != nil {
		log.Fatalf("Failed to load simulConfig: %v", err)
	}
	cli.Cli()
}
