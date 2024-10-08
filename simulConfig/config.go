package simulConfig

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type SimulConfig struct {
	Filepath      string   `json:"filepath" yaml:"filepath"`             // path to the configuration file
	DockerConfigs string   `json:"docker_dir" yaml:"docker_configs"`     // default directory for Docker operations
	ProjectRoot   string   `json:"project_root" yaml:"project_root"`     // default project root
	DockerNetwork string   `json:"docker_network" yaml:"docker_network"` // default Docker network
	Metaservices  []string `json:"metaservices" yaml:"metaservices"`     // default meta services
	// for ssh connection
	SshUser    string `json:"ssh_user" yaml:"ssh_user"`       // default ssh user
	RemoteHost string `json:"remote_host" yaml:"remote_host"` // default remote host
}

var Get = &SimulConfig{
	Filepath:      "",
	DockerConfigs: ".",
	ProjectRoot:   ".",
	DockerNetwork: "",
	Metaservices:  []string{},
	SshUser:       "",
	RemoteHost:    "",
}

func (config *SimulConfig) Hydrate() error {
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

	if err := yaml.NewDecoder(file).Decode(&Get); err != nil {
		return err
	}
	return nil
}

func (config *SimulConfig) Save() error {
	config.ensureFilePath()

	file, err := os.OpenFile(config.Filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return err
	}
	if err := encoder.Close(); err != nil {
		return err
	}

	return nil
}

func (config *SimulConfig) ensureFilePath() {
	if config.Filepath == "" {
		homeDir, _ := os.UserHomeDir()
		config.Filepath = filepath.Join(homeDir, ".simulploy.yaml")
	}
}
