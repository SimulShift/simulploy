package egg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CommandExecutor provides a generalized way to execute system commands using a builder pattern.
type CommandExecutor struct {
	output     *os.File // Standard output for commands
	sudo       bool     // Determines if sudo is used
	scriptPath string   // Path to the script or command to run
	args       []string // Arguments for the command
}

// NewEgg creates a new command executor instance.
func NewEgg(output *os.File) *CommandExecutor {
	return &CommandExecutor{
		output: output,
		args:   []string{},
	}
}

// SetSudo enables running the command with sudo.
func (ce *CommandExecutor) SetSudo() *CommandExecutor {
	ce.sudo = true
	return ce
}

// SetPath sets the path of the script or executable.
func (ce *CommandExecutor) SetPath(path string) *CommandExecutor {
	ce.scriptPath = path
	return ce
}

// AddArg adds an argument to the command.
func (ce *CommandExecutor) AddArg(arg string) *CommandExecutor {
	ce.args = append(ce.args, arg)
	return ce
}

// Run executes the constructed command.
func (ce *CommandExecutor) Run() bool {
	cmd := exec.Command(ce.scriptPath, ce.args...)
	cmd.Stdout = ce.output
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running script: %v for command %s and args: "+
			"%s", err, ce.scriptPath, strings.Join(ce.args, " ")+"\n")
		return false
	}
	return true
}

func (ce *CommandExecutor) String() string {
	return fmt.Sprintf("%v %v", ce.scriptPath, ce.args)
}
