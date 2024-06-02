package docker_parser

import (
	"context"
	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/simulshift/simulploy/egg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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

// LoadComposeFiles loads multiple Docker Compose files and merges them into a single project configuration.
func LoadComposeFiles(files []string) *types.Project {
	// Prepare a slice of ConfigFile structs from the list of filenames.
	options, err := cli.NewProjectOptions(files, cli.WithProfiles(egg.Profiles), cli.WithEnvFiles("./docker/.env"), cli.WithDotEnv)
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
