package cli

import (
	"github.com/simulshift/simulploy/cli/cmd"
	_ "github.com/simulshift/simulploy/cli/cmd/db" // Import the db package to ensure its init() is called
)

func Cli() {
	cmd.Execute()
}
