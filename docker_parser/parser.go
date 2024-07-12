package docker_parser

import (
	"context"
	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/simulshift/simulploy/simulConfig"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DockerCompose represents the root structure of a Docker Compose file, focusing on profiles.
type DockerCompose struct {
	Services map[string]struct {
		Profiles []string `yaml:"profiles"`
	} `yaml:"services"`
}

// GetDockerComposeFiles reads all files in a directory and returns a list of YAML files.
func GetDockerComposeFiles(dockerDir string) []string {
	// Read all files in the directory.
	files, err := os.ReadDir(dockerDir)
	if err != nil {
		log.Fatalf("Failed to read docker directory: %v", err)
	}

	// Filter for YAML files.
	var yamlFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".yml") {
			yamlFiles = append(yamlFiles, filepath.Join(dockerDir, file.Name()))
		}
	}
	return yamlFiles
}

// LoadComposeFilesWithProfiles loads multiple Docker Compose files and merges them into a single project configuration.
func LoadComposeFilesWithProfiles(files []string, profiles []string) *types.Project {
	envPath := filepath.Join(simulConfig.Get.DockerDir, ".env")
	log.Println("envPath: ", envPath)
	options, err := cli.NewProjectOptions(files, cli.WithEnvFiles(envPath), cli.WithProfiles(profiles), cli.WithDotEnv)
	if err != nil {
		log.Fatalf("Failed to create project options: %v", err)
	}
	ctx := context.Background()
	project, err := options.LoadProject(ctx)
	if err != nil {
		log.Fatalf("Failed to load project: %v", err)
	}
	return project
}

// LoadComposeFiles loads all Docker Compose files and merges them into a single project configuration.
// Includes all profiles found in the services.
func LoadComposeFiles(files []string) *types.Project {
	profiles, err := extractProfiles(files)
	if err != nil {
		log.Fatalf("Failed to extract profiles: %v", err)
	}
	return LoadComposeFilesWithProfiles(files, profiles)
}

// extractProfiles parses the YAML content from multiple files to extract all unique profiles from services.
func extractProfiles(files []string) ([]string, error) {
	profileSet := make(map[string]bool)
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var compose DockerCompose
		if err := yaml.Unmarshal(data, &compose); err != nil {
			return nil, err
		}

		for _, service := range compose.Services {
			for _, profile := range service.Profiles {
				profileSet[profile] = true
			}
		}
	}

	profiles := make([]string, 0, len(profileSet))
	for profile := range profileSet {
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
