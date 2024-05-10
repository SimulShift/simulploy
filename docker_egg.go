package main

import (
	"os"
)

type Profile string

const (
	ProfileUnset Profile = "unset"
	Development  Profile = "development"
	Production   Profile = "production"
	Linux        Profile = "linux"
)

type Service string

const (
	ServiceUnset    Service = "unset"
	postgresDev     Service = "postgres-dev"
	postgresProd    Service = "postgres-prod"
	envoyProd       Service = "envoy-prod"
	envoyDevWindows Service = "envoy-dev-windows"
)

// Docker handles Docker Compose operations.
type Docker struct {
	egg   *CommandExecutor
	clean bool
}

//const postgresYaml =

// file paths

// NewDocker creates a manager for Docker services.
func NewDocker() *Docker {
	executor := NewEgg(os.Stdout)
	executor.AddArg("docker")
	executor.AddArg("compose")

	return &Docker{
		egg: executor,
	}
}

// SetClean sets the clean flag for the Docker Compose services.
func (docker *Docker) SetClean(clean bool) *Docker {
	docker.clean = clean
	return docker
}

// Up - starts Docker Compose services with an option to clean them first.
func (docker *Docker) Up() *Docker {
	docker.egg.AddArg("up")
	return docker
}

// Down - stops Docker Compose services.
func (docker *Docker) Down() *Docker {
	docker.egg.AddArg("down")
	return docker
}

func (docker *Docker) Detached() *Docker {
	docker.egg.AddArg("-d")
	return docker
}

// Compose - runs the Docker Compose command.
func (docker *Docker) Compose() {
	// add the filepath with -f
	docker.egg.AddArg("-f")
	docker.egg.AddArg("docker-compose.yml")

	// run the egg
	if !docker.egg.Run() {
		os.Exit(1)
	}

	if docker.clean && docker.egg.args[1] == "down" {
		// run "docker volume rm docker_postgres-data"
		cleanEgg := NewEgg(os.Stdout)
		cleanEgg.AddArg("docker")
		cleanEgg.AddArg("volume")
		cleanEgg.AddArg("rm")
		cleanEgg.AddArg("docker_postgres-data")
		if !cleanEgg.Run() {
			os.Exit(1)
		}
	}
}
