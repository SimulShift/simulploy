package main

import (
	"github.com/joho/godotenv"
	"github.com/simulshift/simulploy/cli"
	"github.com/simulshift/simulploy/simulConfig"
	"log"
	"path/filepath"
)

func main() {
	// try to load the configuration
	err := simulConfig.Get.Hydrate()
	if err != nil {
		log.Fatalf("Failed to load simulConfig: %v", err)
	}
	// Construct the path to the .env file within the project root
	envPath := filepath.Join(simulConfig.Get.ProjectRoot, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cli.Cli()
}
