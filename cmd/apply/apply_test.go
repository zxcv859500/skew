package apply

import (
	"github.com/zxcv859500/skew/pkg/genericclioptions"
	"testing"
)

const (
	filename = "../../testdata/ParseThisYaml.yaml"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		filepath    string
		expectedOut string
	}{
		{
			name:        "apply file",
			args:        []string{},
			filepath:    filename,
			expectedOut: "Job Name: testJob",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ioStreams, _, buf, _ := genericclioptions.NewTestIOStreams()
			apply := NewCmdApply(ioStreams)
			_ = apply.Flags().Set("filename", test.filepath)
			apply.Run(apply, test.args)

			if buf.String() != test.expectedOut {
				t.Fatalf("unexpected error: expected %s to occur, but got %s", test.expectedOut, buf.String())
			}
		})
	}
}
