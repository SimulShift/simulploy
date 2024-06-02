package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ConfigType represents the CLI configuration.
type ConfigType struct {
	Filepath      string `json:"filepath"`       // path to the configuration file
	DockerDir     string `json:"docker_dir"`     // default directory for Docker operations
	DockerNetwork string `json:"docker_network"` // default Docker network
}

var Config = &ConfigType{
	Filepath:  "",
	DockerDir: ".",
}

func (config *ConfigType) Load() error {
	config.ensureFilePath()

	file, err := os.Open(config.Filepath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create and save default configuration if file does not exist
			return config.Save()
		}
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&Config); err != nil {
		return err
	}
	return nil
}

func (config *ConfigType) Save() error {
	config.ensureFilePath()

	file, err := os.OpenFile(config.Filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err = encoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func (config *ConfigType) ensureFilePath() {
	if config.Filepath == "" {
		homeDir, _ := os.UserHomeDir()
		config.Filepath = filepath.Join(homeDir, ".simulploy")
	}
}
