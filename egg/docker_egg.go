package egg

import (
	"fmt"
	"github.com/simulshift/simulploy/docker_parser"
	"github.com/simulshift/simulploy/simulConfig"
	"github.com/simulshift/simulploy/simulSsh"
	"log"
	"os"
	"os/exec"
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
	dockerDir := simulConfig.Get.DockerConfigs
	for _, metaservice := range simulConfig.Get.Metaservices {
		MetaserviceToYaml[metaservice] = filepath.Join(dockerDir, "docker-compose", metaservice+".docker-compose.yaml")
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
	log.Println("Setting metaservice: ", metaservice)
	docker.MetaService = metaservice
	log.Println("Adding yaml for metaservice: ", MetaserviceToYaml[metaservice])
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

	// TODO (alex): Clean up/automate
	if docker.profile == "production" && docker.MetaService == "envoy" && docker.Direction == Up {
		log.Println("Wiring volume folder for envoy")
		docker.WireVolumeFolder(docker.MetaService)
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

func (docker *Docker) WireVolumeFolder(metaservice string) {
	// get config
	simulConfig.Get.Hydrate()
	remoteUser := simulConfig.Get.SshUser
	remoteHost := simulConfig.Get.RemoteHost
	remoteVolumeFolder := "/home/x/volumes/"
	metaserviceFolder := remoteVolumeFolder + "/" + metaservice
	localEnvoyPath := simulConfig.Get.DockerConfigs + "/envoy"
	keyPath := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_ed25519")
	// Use scp to copy the entire folder to the remote server, specifying the private key
	// first make sure directory exists
	sshClient := simulSsh.New() // Create SSH connection
	defer sshClient.Close()

	// Step 1: Remove the remote directory if it exists
	deleteCmd := fmt.Sprintf("rm -rf %s", metaserviceFolder)
	_, err := sshClient.Exec(deleteCmd) // Use SSH client to execute
	if err != nil {
		log.Fatalf("Failed to remove remote directory: %v", err)
	}

	err = sshClient.EnsureRemoteDirExists(metaserviceFolder)
	if err != nil {
		log.Fatalf("Failed to create remote directory: %v", err)
	}

	cmd := exec.Command("scp", "-r", "-i", keyPath,
		localEnvoyPath, // source directory
		fmt.Sprintf("%s@%s:%s", remoteUser, remoteHost, remoteVolumeFolder))

	// Capture the output of scp for logging purposes
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to scp envoy folder: %v\nOutput: %s", err, string(output))
	} else {
		log.Printf("Successfully copied envoy folder. Output: %s", string(output))
	}

	// Apply dos2unix to all files in the envoy directory
	err = sshClient.ApplyDos2UnixRemote(remoteVolumeFolder)
	if err != nil {
		log.Fatalf("Failed to run dos2unix on all files: %v", err)
	}
}

func (docker *Docker) DropDatabase() {
	// only if the profile is postgres
	if docker.profile != "postgres" {
		log.Println("Profile is not postgres")
		return
	}
	// run "docker volume rm docker_postgres-data"
	cleanEgg := NewEgg(os.Stdout).SetSudo().SetPath("docker").
		AddArg("volume").AddArg("rm").AddArg("docker_postgres-data")
	if !cleanEgg.Run() {
		os.Exit(1)
	}
}
