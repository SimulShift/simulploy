package simulSsh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"log"
	"os"
	"path/filepath"
)

type simulSsh struct {
	sshServerIp string
	sshUsername string
	Config      *ssh.ClientConfig
	sshClient   *ssh.Client
}

func New() *simulSsh {
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
	sshUser := os.Getenv("SSH_USER")
	if sshUser == "" {
		log.Fatalf("SSH_USER not set")
	}

	// get ip from .env file
	sshServerIp := os.Getenv("SSH_IP")
	if sshServerIp == "" {
		log.Fatalf("SSH_IP not set")
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

	return &simulSsh{
		Config:      config,
		sshServerIp: sshServerIp,
		sshUsername: sshUser,
		sshClient:   client,
	}
}

func (s *simulSsh) Exec(command string) (string, error) {
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

func (s *simulSsh) Close() error {
	if s.sshClient != nil {
		return s.sshClient.Close()
	}
	return nil
}
