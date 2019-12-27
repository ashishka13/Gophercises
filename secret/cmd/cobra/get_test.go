package cobra

import (
	"testing"

	"github.com/spf13/cobra"
)

//TestGetErr is used to check whether
// invalid key is giving error or not
func TestGetErr(t *testing.T) {
	var cmd *cobra.Command
	getCmd.Run(cmd, []string{"this kay is not present"})
}

func TestGet(t *testing.T) {
	var cmd *cobra.Command
	getCmd.Run(cmd, []string{"demo"})
}
