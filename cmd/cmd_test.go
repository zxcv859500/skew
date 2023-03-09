package cmd_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zxcv859500/skew/cmd"
)

func TestNewDefaultSkewCommand(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		expect string
	}{
		{
			name:   "test without warnings-as-errors flag",
			args:   []string{},
			expect: "A CLI tool to extract web page elements using YAML config files\n\nUsage:\n  skew [command]\n\nAvailable Commands:\n  apply       Long desc\n  completion  Generate the autocompletion script for the specified shell\n  help        Help about any command\n\nFlags:\n  -h, --help                 help for skew\n      --warnings-as-errors   Treat warnings received from the server as errors and exit with a non-zero exit code\n\nUse \"skew [command] --help\" for more information about a command.\n",
		},
		{
			name:   "test with warnings-as-errors flag",
			args:   []string{"--warnings-as-errors"},
			expect: "A CLI tool to extract web page elements using YAML config files\n\nUsage:\n  skew [command]\n\nAvailable Commands:\n  apply       Long desc\n  completion  Generate the autocompletion script for the specified shell\n  help        Help about any command\n\nFlags:\n  -h, --help                 help for skew\n      --warnings-as-errors   Treat warnings received from the server as errors and exit with a non-zero exit code\n\nUse \"skew [command] --help\" for more information about a command.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a new default Skew command
			cmd := cmd.NewDefaultSkewCommand()

			// capture the output of the command
			var buf bytes.Buffer
			cmd.SetOutput(&buf)

			// set the command arguments
			cmd.SetArgs(tt.args)

			// execute the command
			err := cmd.Execute()

			// assert that there was no error
			assert.NoError(t, err)

			// assert that the command output matches the expected output
			assert.Equal(t, tt.expect, buf.String())
		})
	}
}
