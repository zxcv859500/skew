package parser

import (
	"context"
	"github.com/zxcv859500/skew/pkg/resource"
	"github.com/zxcv859500/skew/workflow"
	"gopkg.in/yaml.v3"
	"os"
)

type errTimeoutError struct{}

func (errTimeoutError) Error() string {
	return "timeout while read file."
}

func ParseYaml(buf []byte) (*workflow.Workflow, error) {
	session := &workflow.Workflow{}
	err := yaml.Unmarshal(buf, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func ReadYamlFile(ctx context.Context, options resource.FilenameOptions) (*workflow.Workflow, error) {
	done := make(chan []byte, 1)
	errs := make(chan error, 1)
	go func() {
		buf, err := os.ReadFile(options.Filenames[0])
		if err != nil {
			errs <- err
		}
		done <- buf
	}()

	select {
	case err := <-errs:
		return nil, err
	case <-ctx.Done():
		return nil, errTimeoutError{}
	case buf := <-done:
		return ParseYaml(buf)
	}
}
