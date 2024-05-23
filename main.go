package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/simulshift/simulploy/egg"
	"log"
	"os"
	"slices"
)

// write main function
func main() {
	DockerCli()

	/*
		// start the server
		go server.StartServer()
		// start the client
		log.Println("Starting the client")
		greeter_client.GreeterClient()
	*/
}

func DockerCli() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// check if --down flag is provided
	downFlag := flag.Bool("down", false, "to docker compose down")
	// clean flag
	cleanFlag := flag.Bool("clean", false, "to clean the docker images")
	profileFlag := flag.String("profile", "development", "profile to use")
	flag.Parse()
	// create a new Docker instance
	docker := egg.NewDocker()
	// Validate the provided profile
	if !slices.Contains(egg.ValidProfiles, egg.Profile(*profileFlag)) {
		fmt.Println("Invalid profile provided")
		os.Exit(1)
	}
	docker.SetProfile(egg.Profile(*profileFlag))
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <profile> or <profile, metaservice")
		os.Exit(1)
	}
	if len(args) > 0 {
		// Get the metaservice from the command line
		metaservice := egg.MetaService(args[1])
		// Check if metaservice is key in the map
		if _, ok := egg.MetaserviceToYaml[metaservice]; !ok {
			fmt.Println("Invalid metaservice provided")
			os.Exit(1)
		}
		docker.SetMetaService(metaservice)
	}
	if *downFlag {
		log.Println("Docker compose down")
		docker.Down()
	} else {
		docker.Up()
	}

	if *cleanFlag {
		docker.Clean()
	}

	docker.Compose()
}
