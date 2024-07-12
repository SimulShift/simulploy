package egg

import (
	"github.com/simulshift/simulploy/docker_parser"
	"github.com/simulshift/simulploy/simulConfig"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var MetaserviceToYaml = make(map[string]string)
var Profiles []string

type Direction string

const (
	DirectionUnset Direction = "unset"
	Up             Direction = "up"
	Down           Direction = "down"
)

// Docker handles Docker Compose operations.
type Docker struct {
	egg         *Egg
	profile     string
	MetaService string
	Direction   Direction
	clean       bool
	drop        bool
	detached    bool
}

// getAllKeyValues - gets all the keys and values from a map
func getAllKeyValues(m map[string]string) ([]string, []string) {
	var values []string
	var keys []string
	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

// GetProfiles - gets the profiles defined in the yaml file
func GetProfiles() []string {
	// get list of yaml files from map
	_, yamlFiles := getAllKeyValues(MetaserviceToYaml)
	project := docker_parser.LoadComposeFiles(yamlFiles)
	Profiles = project.Profiles
	return Profiles
}

// NewDocker creates a manager for Docker services.
func NewDocker() *Docker {
	simulConfig.Get.Hydrate()
	dockerDir := simulConfig.Get.DockerDir
	for _, metaservice := range simulConfig.Get.Metaservices {
		MetaserviceToYaml[metaservice] = filepath.Join(dockerDir, metaservice+".docker-compose.yaml")
	}
	GetProfiles()
	egg := NewEgg(os.Stdout)
	egg.AddArg("compose")
	return &Docker{
		egg:         egg,
		profile:     "",
		MetaService: "",
		Direction:   DirectionUnset,
		clean:       false,
		drop:        false,
		detached:    false,
	}
}

// AddYamlIfNotAlreadyAdded adds the yaml file to the commands if it's not already added.
func (docker *Docker) addYamlIfNotAlreadyAdded(yaml string) {
	if !slices.Contains(docker.egg.args, yaml) {
		docker.AddDockerComposeFile(yaml)
	}
}

func (docker *Docker) SetMetaService(metaservice string) *Docker {
	if metaservice == "" {
		log.Println("Empty metaservice provided")
		os.Exit(1)
	}
	// environment must first be set
	if docker.profile == "" {
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

func (docker *Docker) SetProfile(profile string) *Docker {
	// set the profile
	// validate the profile
	if !slices.Contains(Profiles, profile) {
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
	docker.detached = true
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
	if docker.profile == "" {
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

	if docker.detached {
		docker.egg.AddArg("-d")
	}

	// run the egg
	if !docker.egg.Run() {
		// log error
		log.Println("Error running Docker Compose for: " + docker.egg.String())
		os.Exit(1)
	}

	if docker.drop {
		docker.DropDatabase()
	}
}

func (docker *Docker) DropDatabase() {
	// run "docker volume rm docker_postgres-data"
	cleanEgg := NewEgg(os.Stdout).SetSudo().SetPath("docker").
		AddArg("volume").AddArg("rm").AddArg("docker_postgres-data")
	if !cleanEgg.Run() {
		os.Exit(1)
	}
}
