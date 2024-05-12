package main

import "github.com/simulshift/simulploy/egg"

// write main function
func main() {
	// create a new Docker instance
	docker := egg.NewDocker()
	// start the Docker services
	docker.AddDockerComposeFile("docker/docker-compose.postgres.yaml").Down().SetClean(true).Compose()
}
