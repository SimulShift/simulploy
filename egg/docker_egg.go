package egg

import (
	"log"
	"os"
	"slices"
)

var debug = false

var PostgresYaml = "docker/docker-compose.postgres.yaml"
var EnvoyYaml = "docker/docker-compose.envoy.yaml"
var ChatbotYaml = "docker/docker-compose.chatbot.yaml"

// ServiceYamlMap map of services to their respective docker-compose files
var ServiceYamlMap = map[Service]string{
	PostgresDev:     PostgresYaml,
	PostgresProd:    PostgresYaml,
	EnvoyDevWindows: EnvoyYaml,
	EnvoyProd:       EnvoyYaml,
	ChatbotProd:     ChatbotYaml,
}

// MetaserviceToYaml Map meta services to file
var MetaserviceToYaml = map[MetaService]string{
	Postgres: PostgresYaml,
	Envoy:    EnvoyYaml,
	Chatbot:  ChatbotYaml,
}

type MetaService string

const (
	Unset    MetaService = "unset"
	Postgres MetaService = "postgres"
	Envoy    MetaService = "envoy"
	Chatbot  MetaService = "chatbot"
)

type Profile string

const (
	ProfileUnset Profile = "unset"
	Development  Profile = "development"
	Production   Profile = "production"
	Linux        Profile = "linux"
)

// ValidProfiles Define an array of valid profiles
var ValidProfiles = []Profile{
	ProfileUnset, // You can include this if it's a valid profile to use
	Development,
	Production,
	Linux,
}

type Service string

const (
	ServiceUnset    Service = "unset"
	PostgresDev     Service = "postgres-dev"
	PostgresProd    Service = "postgres-prod"
	EnvoyProd       Service = "envoy-prod"
	EnvoyDevWindows Service = "envoy-dev-windows"
	ChatbotProd     Service = "chatbot-prod"
)

// ValidServices Define an array of valid services
var ValidServices = []Service{
	ServiceUnset, // You can include this if it's a valid services to use
	PostgresDev,
	PostgresProd,
	EnvoyProd,
	EnvoyDevWindows,
}

type Direction string

const (
	DirectionUnset Direction = "unset"
	Up             Direction = "up"
	Down           Direction = "down"
)

// Docker handles Docker Compose operations.
type Docker struct {
	egg         *Egg
	services    []Service
	profile     Profile
	MetaService MetaService
	Direction   Direction
	clean       bool
}

// NewDocker creates a manager for Docker services.
func NewDocker() *Docker {
	executor := NewEgg(os.Stdout)
	executor.AddArg("compose")

	return &Docker{
		egg:         executor,
		services:    make([]Service, 0), // empty slice
		profile:     ProfileUnset,
		MetaService: Unset,
		Direction:   DirectionUnset,
		clean:       false,
	}
}

// AddYamlIfNotAlreadyAdded adds the yaml file to the commands if it's not already added.
func (docker *Docker) addYamlIfNotAlreadyAdded(yaml string) {
	if !slices.Contains(docker.egg.args, yaml) {
		docker.AddDockerComposeFile(yaml)
	}
}

func (docker *Docker) SetMetaService(metaservice MetaService) *Docker {
	// environment must first be set
	if docker.profile == ProfileUnset {
		log.Println("Profile must be set before setting metaservice")
		os.Exit(1)
	}
	docker.addYamlIfNotAlreadyAdded(MetaserviceToYaml[metaservice])
	return docker
}

func (docker *Docker) AddService(service Service) *Docker {
	docker.services = append(docker.services, service)
	// get the docker-compose file for the services
	if yaml, ok := ServiceYamlMap[service]; ok {
		docker.addYamlIfNotAlreadyAdded(yaml)
	}
	return docker
}

func (docker *Docker) SetProfile(profile Profile) *Docker {
	docker.profile = profile
	return docker
}

// SetClean sets the clean flag for the Docker Compose services.
func (docker *Docker) SetClean(clean bool) *Docker {
	docker.clean = clean
	return docker
}

// Up - starts Docker Compose services with an option to clean them first.
func (docker *Docker) Up() *Docker {
	docker.Direction = Up
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

func (docker *Docker) AddDockerComposeFile(filepath string) *Docker {
	docker.egg.AddArg("-f")
	docker.egg.AddArg(filepath)
	return docker
}

// Compose - runs the Docker Compose command.
func (docker *Docker) Compose() {
	docker.egg.SetPath("docker")

	// add services
	for _, service := range docker.services {
		docker.egg.AddArg(string(service))
	}

	// set profile
	if docker.profile == ProfileUnset {
		log.Println("Profile must be set before running Docker Compose")
		os.Exit(1)
	}

	// add the profile to the command
	docker.egg.AddArg("--profile")
	docker.egg.AddArg(string(docker.profile))

	// add the direction to the command
	if docker.Direction != DirectionUnset {
		docker.egg.AddArg(string(docker.Direction))
	}

	if true {
		log.Println("Running Docker Compose for: " + docker.egg.String())
	}

	// run the egg
	if !docker.egg.Run() {
		// log error
		log.Println("Error running Docker Compose for: " + docker.egg.String())
		os.Exit(1)
	}

	/* TODO: Add clean up for volumes
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
	*/
}
