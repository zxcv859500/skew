package resource

import "fmt"

var FileExtensions = []string{".json", ".yaml", ".yml"}
var InputExtensions = append(FileExtensions, "stdin")

const defaultHttpGetAttempts = 3
const pathNotExistError = "the path %q does not exist"

type FilenameOptions struct {
	Filenames []string
	Recursive bool
}

func (o *FilenameOptions) RequireFilename() error {
	if len(o.Filenames) == 0 {
		return fmt.Errorf("must specify -f")
	}
	return nil
}
