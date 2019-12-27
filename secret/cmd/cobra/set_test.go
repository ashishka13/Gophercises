package cobra

import (
	"testing"

	"github.com/spf13/cobra"
)

func change() {
	Fake = 1
}

// TestGet checks for valid key
func TestSet(t *testing.T) {
	var cmd *cobra.Command
	args := []string{"", ""}
	setCmd.Run(cmd, args)
}

func TestSetError(t *testing.T) {
	var cmd *cobra.Command
	args := []string{"", ""}
	change()
	setCmd.Run(cmd, args)
}
