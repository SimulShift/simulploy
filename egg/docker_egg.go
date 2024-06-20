package egg

import (
	"log"
	"os"
	"path/filepath"
	"slices"
)

// BasePath imported via ldflags
var PostgresYaml = "docker/docker-compose.postgres.yaml"
var EnvoyYaml = "docker/docker-compose.envoy.yaml"
var ChatbotYaml = "docker/docker-compose.chatbot.yaml"

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

// ValidMetaServices Valid MetaServices
var ValidMetaServices = []MetaService{
	Unset,
	Postgres,
	Envoy,
	Chatbot,
}

type Profile string

const (
	ProfileUnset Profile = "unset"
	Development  Profile = "development"
	Production   Profile = "production"
	Linux        Profile = "linux"
)

// ValidProfiles Define an array of valid profiles
var ValidProfiles = []Profile{
	ProfileUnset,
	Development,
	Production,
	Linux,
}

var Profiles = []string{
	"development",
	"production",
	"linux",
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
	profile     Profile
	MetaService MetaService
	Direction   Direction
	clean       bool
	drop        bool
}

// NewDocker creates a manager for Docker services.
func NewDocker(dockerDir string) *Docker {
	PostgresYaml = filepath.Join(dockerDir, "docker-compose.postgres.yaml")
	EnvoyYaml = filepath.Join(dockerDir, "docker-compose.envoy.yaml")
	ChatbotYaml = filepath.Join(dockerDir, "docker-compose.chatbot.yaml")
	MetaserviceToYaml = map[MetaService]string{
		Postgres: PostgresYaml,
		Envoy:    EnvoyYaml,
		Chatbot:  ChatbotYaml,
	}

	egg := NewEgg(os.Stdout)
	egg.AddArg("compose")

	return &Docker{
		egg:         egg,
		profile:     ProfileUnset,
		MetaService: Unset,
		Direction:   DirectionUnset,
		clean:       false,
		drop:        false,
	}
}

// AddYamlIfNotAlreadyAdded adds the yaml file to the commands if it's not already added.
func (docker *Docker) addYamlIfNotAlreadyAdded(yaml string) {
	if !slices.Contains(docker.egg.args, yaml) {
		docker.AddDockerComposeFile(yaml)
	}
}

func (docker *Docker) SetMetaService(metaservice MetaService) *Docker {
	if metaservice == "" {
		log.Println("Empty metaservice provided")
		os.Exit(1)
	}
	// environment must first be set
	if docker.profile == ProfileUnset {
		log.Println("Profile must be set before setting metaservice")
		os.Exit(1)
	}
	if MetaserviceToYaml[metaservice] == "" {
		log.Println("Invalid metaservice provided")
		os.Exit(1)
	}
	log.Println("Setting metaservice: ", MetaserviceToYaml[metaservice])
	docker.addYamlIfNotAlreadyAdded(MetaserviceToYaml[metaservice])
	return docker
}

func (docker *Docker) SetProfile(profile Profile) *Docker {
	// set the profile
	// validate the profile
	if !slices.Contains(ValidProfiles, profile) {
		log.Println("Invalid profile provided")
		os.Exit(1)
	}
	docker.profile = profile
	return docker
}

// Clean sets the clean flag for the Docker Compose services.
func (docker *Docker) Clean() *Docker {
	docker.Down()
	docker.clean = true
	return docker
}

// Drop sets the drop flag for the Docker Compose services.
func (docker *Docker) Drop() *Docker {
	docker.drop = true
	return docker
}

// Up - starts Docker Compose services with an option to clean them first.
func (docker *Docker) Up() *Docker {
	docker.Direction = Up
	return docker
}

// Down - stops Docker Compose services.
func (docker *Docker) Down() *Docker {
	docker.Direction = Down
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

	log.Println("Running Docker Compose for: " + docker.egg.String())

	if docker.clean {
		docker.egg.AddArg("--rmi")
		docker.egg.AddArg("all")
	}

	if docker.drop {
		docker.DropDatabase()
	}

	// run the egg
	if !docker.egg.Run() {
		// log error
		log.Println("Error running Docker Compose for: " + docker.egg.String())
		os.Exit(1)
	}

}

func (docker *Docker) DropDatabase() {
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
