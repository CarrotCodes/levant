package command

import (
	"bytes"
	"fmt"

	"github.com/jrasell/levant/buildtime"
	"github.com/mitchellh/cli"
)

// VersionCommand is a Command implementation that prints the version.
type VersionCommand struct {
	Revision          string
	Version           string
	VersionPrerelease string
	UI                cli.Ui
}

// Help provides the help information for the version command.
func (c *VersionCommand) Help() string {
	return ""
}

// Synopsis is provides a brief summary of the version command.
func (c *VersionCommand) Synopsis() string {
	return "Prints the Levant version"
}

// Run executes the version command.
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "Levant v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.UI.Output(versionString.String())
	c.UI.Output(fmt.Sprintf("Date: %s", buildtime.BuildDate))
	c.UI.Output(fmt.Sprintf("Commit: %s", buildtime.GitCommit))
	c.UI.Output(fmt.Sprintf("Branch: %s", buildtime.GitBranch))
	c.UI.Output(fmt.Sprintf("State: %s", buildtime.GitState))
	c.UI.Output(fmt.Sprintf("Summary: %s", buildtime.GitSummary))
	return 0
}
