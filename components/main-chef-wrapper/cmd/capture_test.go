package cmd

import (
	"log"
	"testing"

	"github.com/chef/chef-workstation/components/main-chef-wrapper/dist"
)

func TestCaptureCommand(t *testing.T) {
	for _, test := range []struct {
		Args []string
	}{
		{
			Args: []string{"capture", "--help"},
		},
		{
			Args: []string{"capture", "NODE-NAME"},
		},
	} {
		t.Run("", func(t *testing.T) {
			err := Runner.passThroughCommand(dist.WorkstationExec, "", test.Args)
			if err != nil {
				log.Printf("Command finished with error: %v", err)
			} else {
				log.Printf("Command executed successfully")
			}
		})
	}
}
