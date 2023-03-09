package parser

import (
	"context"
	"github.com/zxcv859500/skew/pkg/resource"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte(`session:
  job:
    - name: testJob
      action: click
      with:
        xpath: testxpath
        test: test
`)
	workflow, err := ParseYaml(input)

	if err != nil {
		panic(err)
	}

	if workflow.Sess.Jobs[0].JobName != "testJob" ||
		workflow.Sess.Jobs[0].Action != "click" {
		panic("error")
	}
}

func TestParseFile(t *testing.T) {
	inputFileName := "ParseThisYaml.yaml"
	inputFilePath := filepath.Join("../testdata", inputFileName)

	workflow, err := ReadYamlFile(context.TODO(), resource.FilenameOptions{
		Filenames: []string{inputFilePath},
		Recursive: false,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if workflow.Sess.Jobs[0].JobName != "testJob" {
		t.Fatalf("unexpected error: expect job name testJob but got %s", workflow.Sess.Jobs[0].JobName)
	}
}
