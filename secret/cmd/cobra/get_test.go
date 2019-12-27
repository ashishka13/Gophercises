package cobra

import (
	"testing"

	"github.com/spf13/cobra"
)

//invalid key provided
func TestGetErr(t *testing.T) {
	var cmd *cobra.Command
	getCmd.Run(cmd, []string{"this kay is not present"})
}

//positive test
func TestGet(t *testing.T) {
	var cmd *cobra.Command
	getCmd.Run(cmd, []string{"demo"})
}
