package main

import (
	"fmt"
	"github.com/simulshift/simulploy/egg"
	"os"
	"slices"
)

var Debug = false

// write main function
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <profile> or <profile, metaservice")
		os.Exit(1)
	}
	// create a new Docker instance
	docker := egg.NewDocker()

	// Get the environment from the command line
	profile := egg.Profile(os.Args[1])
	// Validate the provided profile
	if !slices.Contains(egg.ValidProfiles, profile) {
		fmt.Println("Invalid profile provided")
		os.Exit(1)
	}
	docker.SetProfile(profile)

	// check if 2nd argument is provided
	if len(os.Args) > 2 {
		// Get the metaservice from the command line
		metaservice := egg.MetaService(os.Args[2])
		// Check if metaservice is key in the map
		if _, ok := egg.MetaserviceToYaml[metaservice]; !ok {
			fmt.Println("Invalid metaservice provided")
			os.Exit(1)
		}
		docker.SetMetaService(metaservice)
	}

	// start the Docker services
	docker.Up().Compose()
}
