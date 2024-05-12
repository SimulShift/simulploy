package egg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Egg provides a generalized way to execute system commands using a builder pattern.
type Egg struct {
	output     *os.File // Standard output for commands
	sudo       bool     // Determines if sudo is used
	scriptPath string   // Path to the script or command to run
	args       []string // Arguments for the command
}

// NewEgg creates a new command executor instance.
func NewEgg(output *os.File) *Egg {
	return &Egg{
		output: output,
		args:   []string{},
	}
}

// SetSudo enables running the command with sudo.
func (egg *Egg) SetSudo() *Egg {
	egg.sudo = true
	return egg
}

// SetPath sets the path of the script or executable.
func (egg *Egg) SetPath(path string) *Egg {
	egg.scriptPath = path
	return egg
}

// AddArg adds an argument to the command.
func (egg *Egg) AddArg(arg string) *Egg {
	egg.args = append(egg.args, arg)
	return egg
}

// Run executes the constructed command.
func (egg *Egg) Run() bool {
	cmd := exec.Command(egg.scriptPath, egg.args...)
	cmd.Stdout = egg.output
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running script: %v for command %s and args: "+
			"%s", err, egg.scriptPath, " '"+strings.Join(egg.args, " ")+"'\n")
		return false
	}
	return true
}

func (egg *Egg) String() string {
	return fmt.Sprintf("%v %v", egg.scriptPath, egg.args)
}
