package cobra

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

type Vault struct {
	err error
}

func myFile(encodingKey, filepath string) *Vault {
	return &Vault{
		err: errors.New(""),
	}
}

// TestGet checks for valid key
func TestSet(t *testing.T) {
	var cmd *cobra.Command
	setCmd.Run(cmd, []string{"demo", "abc123"})
}

// // //TestSetErr is mocked test for set command
// func TestSetErr(t *testing.T) {
// 	f := &Vault{err: errors.New("User defined error cipher")}
// 	cobra.FakeSet = f.myFile

// 	var cmd *cobra.Command
// 	setCmd.Run(cmd, []string{"demo", "abc123"})
// }
