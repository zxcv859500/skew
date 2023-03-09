package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zxcv859500/skew/cmd/apply"
	"github.com/zxcv859500/skew/pkg/genericclioptions"
	"github.com/zxcv859500/skew/pkg/rest"
	"os"
)

type SkewOptions struct {
	Arguments       []string
	WarningAsErrors bool

	genericclioptions.IOStreams
}

func NewDefaultSkewCommand() *cobra.Command {
	return NewDefaultSkewCommandWithArgs(SkewOptions{
		Arguments:       os.Args,
		WarningAsErrors: false,
		IOStreams:       genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr},
	})
}

func NewDefaultSkewCommandWithArgs(o SkewOptions) *cobra.Command {
	cmd := NewSkewCommand(o)

	return cmd
}

func NewSkewCommand(o SkewOptions) *cobra.Command {
	warningHandler := rest.NewWarningWriter(o.IOStreams.ErrOut, rest.WarningWriterOptions{Deduplicate: true, Color: true})

	cmds := &cobra.Command{
		Use:   "skew",
		Short: "skew get elements on web page by read yaml file",
		Long:  "A CLI tool to extract web page elements using YAML config files",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			rest.SetDefaultWarningHandler(warningHandler)

			return nil
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			if o.WarningAsErrors {
				count := warningHandler.WarningCount()
				switch count {
				case 0:
					//no warnings
				case 1:
					return fmt.Errorf("%d warning received", count)
				default:
					return fmt.Errorf("%d warnings received", count)
				}
			}
			return nil
		},
	}

	flags := cmds.PersistentFlags()

	flags.BoolVar(&o.WarningAsErrors, "warnings-as-errors", o.WarningAsErrors, "Treat warnings received from the server as errors and exit with a non-zero exit code")

	applyCmd := apply.NewCmdApply(o.IOStreams)

	cmds.AddCommand(applyCmd)

	return cmds
}
