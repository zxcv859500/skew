package apply

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zxcv859500/skew/cmd/utils"
	"github.com/zxcv859500/skew/parser"
	"github.com/zxcv859500/skew/pkg/genericclioptions"
	"github.com/zxcv859500/skew/pkg/resource"
)

type ApplyFlags struct {
	FileNameFlags *genericclioptions.FileNameFlags

	genericclioptions.IOStreams
}

type ApplyOptions struct {
	resource.FilenameOptions

	genericclioptions.IOStreams
}

var (
	applyLong    = "Long desc"
	applyExample = "Example desc"
)

func NewApplyFlags(streams genericclioptions.IOStreams) *ApplyFlags {
	var filenames []string
	recursive := false

	return &ApplyFlags{
		FileNameFlags: &genericclioptions.FileNameFlags{Filenames: &filenames, Recursive: &recursive},
		IOStreams:     streams,
	}
}

func NewCmdApply(ioStreams genericclioptions.IOStreams) *cobra.Command {
	flags := NewApplyFlags(ioStreams)

	cmd := &cobra.Command{
		Use:                   "apply (-f FILENAME | -k DIRECTORY)",
		DisableFlagsInUseLine: true,
		Short:                 applyLong,
		Example:               applyExample,
		Run: func(cmd *cobra.Command, args []string) {
			o, err := flags.ToOptions(cmd, args)
			utils.CheckErr(err)
			utils.CheckErr(o.Validate())
			utils.CheckErr(o.Run())
		},
	}

	flags.AddFlags(cmd)

	return cmd
}

func (o *ApplyOptions) Validate() error {
	err := o.RequireFilename()
	return err
}

func (o *ApplyOptions) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	workflow, err := parser.ReadYamlFile(ctx, o.FilenameOptions)

	_, _ = fmt.Fprintf(o.IOStreams.Out, "Job Name: %s", workflow.Sess.Jobs[0].JobName)

	return err
}

func (flags *ApplyFlags) ToOptions(cmd *cobra.Command, args []string) (*ApplyOptions, error) {
	if len(args) != 0 {
		return nil, utils.UsageErrorf(cmd, "Unexpected args: %v", args)
	}

	o := &ApplyOptions{
		FilenameOptions: flags.FileNameFlags.ToOptions(),
		IOStreams:       flags.IOStreams,
	}

	return o, nil
}

func (flags *ApplyFlags) AddFlags(cmd *cobra.Command) {
	flags.FileNameFlags.AddFlags(cmd.Flags())
}
