package genericclioptions

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/zxcv859500/skew/pkg/resource"
	"strings"
)

type FileNameFlags struct {
	Usage string

	Filenames *[]string
	Recursive *bool
}

func (o *FileNameFlags) ToOptions() resource.FilenameOptions {
	options := resource.FilenameOptions{}
	if o == nil {
		return options
	}

	if o.Recursive != nil {
		options.Recursive = *o.Recursive
	}

	if o.Filenames != nil {
		options.Filenames = *o.Filenames
	}

	return options
}

func (o *FileNameFlags) AddFlags(flags *pflag.FlagSet) {
	if o == nil {
		return
	}

	if o.Recursive != nil {
		flags.BoolVarP(o.Recursive, "recursive", "R", *o.Recursive, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	}

	if o.Filenames != nil {
		flags.StringSliceVarP(o.Filenames, "filename", "f", *o.Filenames, o.Usage)
		annotations := make([]string, 0, len(resource.FileExtensions))
		for _, ext := range resource.FileExtensions {
			annotations = append(annotations, strings.TrimLeft(ext, "."))
		}
		_ = flags.SetAnnotation("filename", cobra.BashCompFilenameExt, annotations)
	}
}
