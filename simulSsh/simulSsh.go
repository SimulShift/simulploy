package simulSsh

import (
	"bytes"
	"fmt"
	"github.com/simulshift/simulploy/simulConfig"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SimulSsh struct {
	sshServerIp string
	sshUsername string
	Config      *ssh.ClientConfig
	sshClient   *ssh.Client
}

func New() *SimulSsh {
	simulConfig.Get.Hydrate()
	// Path to the private key
	keyPath := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_ed25519")

	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	// Path to the known_hosts file
	knownHostsPath := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "known_hosts")
	hostKeyCallback, err := knownhosts.New(knownHostsPath)
	if err != nil {
		log.Fatalf("failed to create host key callback function: %v", err)
	}

	// get user from .env file
	sshUser := simulConfig.Get.SshUser
	if sshUser == "" {
		log.Fatalf("SSH_USER not set")
	}

	// get ip from .env file
	sshServerIp := simulConfig.Get.RemoteHost + ":22"
	if sshServerIp == "" {
		log.Fatalf("REMOTE_HOST not set")
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	client, err := ssh.Dial("tcp", sshServerIp, config)
	if err != nil {
		log.Fatalf("failed to dial SSH server: %v", err)
	}

	return &SimulSsh{
		Config:      config,
		sshServerIp: sshServerIp,
		sshUsername: sshUser,
		sshClient:   client,
	}
}

func (s *SimulSsh) Exec(command string) (string, error) {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return "Error occurred", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	if err := session.Run(command); err != nil {
		return "Error occurred", fmt.Errorf("failed to run command: %w", err)
	}

	return stdoutBuf.String(), nil
}

func (s *SimulSsh) Close() error {
	if s.sshClient != nil {
		return s.sshClient.Close()
	}
	return nil
}

// EnsureRemoteDirExists - Function to ensure remote directory exists
func (s *SimulSsh) EnsureRemoteDirExists(remoteDir string) error {
	command := fmt.Sprintf("mkdir -p %s", remoteDir)
	output, err := s.Exec(command)
	if err != nil {
		return fmt.Errorf("failed to create remote directory: %w, output: %s", err, output)
	}
	log.Printf("Remote directory created or already exists: %s", remoteDir)
	return nil
}

// ApplyDos2UnixRemote applies dos2unix to every file in the given directory recursively on a remote server
func (s *SimulSsh) ApplyDos2UnixRemote(remoteDirPath string) error {
	// Use find command to get all regular files in the directory
	command := fmt.Sprintf("find %s -type f", remoteDirPath)
	fileList, err := s.Exec(command) // Execute the find command on the remote server
	if err != nil {
		return fmt.Errorf("failed to list files in remote directory %s: %v", remoteDirPath, err)
	}

	// Split the file list by newline to process each file individually
	files := strings.Split(fileList, "\n")
	for _, file := range files {
		if file != "" {
			log.Printf("Running dos2unix on remote file: %s", file)

			// Run dos2unix on each file via SSH
			dos2unixCmd := fmt.Sprintf("dos2unix %s", file)
			output, err := s.Exec(dos2unixCmd)
			if err != nil {
				return fmt.Errorf("failed to run dos2unix on %s: %v\nOutput: %s", file, err, output)
			}
		}
	}
	return nil
}
