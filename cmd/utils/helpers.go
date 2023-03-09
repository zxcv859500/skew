package utils

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const (
	ApplyAnnotationsFlag = "save-config"
	DefaultErrorExitCode = 1
	DefaultChunkSize     = 500
)

type debugError interface {
	DebugError() (msg string, args []interface{})
}

func AddSourceToErr(verb string, source string, err error) error {
	if source != "" {
		return fmt.Errorf("error when %s %q: %v", verb, source, err)
	}
	return err
}

var fatalErrHandler = fatal

func BehaviorOnFatal(f func(string, int)) {
	fatalErrHandler = f
}

func DefaultBehaviorOnFatal() {
	fatalErrHandler = fatal
}

func fatal(msg string, code int) {
	if glog.V(99) {
		glog.FatalDepth(2, msg)
	}
	if len(msg) > 0 {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		_, _ = fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

var ErrExit = fmt.Errorf("exit")

func CheckErr(err error) {
	checkErr(err, fatalErrHandler)
}

func CheckDiffErr(err error) {
	checkErr(err, func(msg string, code int) {
		fatalErrHandler(msg, code+1)
	})
}

func checkErr(err error, handleErr func(string, int)) {
	if err == nil {
		return
	}

	switch {
	case err == ErrExit:
		handleErr("", DefaultErrorExitCode)
	default:
		switch err := err.(type) {
		default:
			msg, ok := StandardErrorMessage(err)
			if !ok {
				msg = err.Error()
				if !strings.HasPrefix(msg, "error: ") {
					msg = fmt.Sprintf("error: %s", msg)
				}
			}
			handleErr(msg, DefaultErrorExitCode)
		}
	}
}

func StandardErrorMessage(err error) (string, bool) {
	if debugErr, ok := err.(debugError); ok {
		glog.V(4).Infof(debugErr.DebugError())
	}
	return "", false
}

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '$%s -h' for help and examples", msg, cmd.CommandPath())
}

func ValidateFlag(cmd *cobra.Command, flag string, expectValue interface{}) (interface{}, bool) {
	switch expectValue.(type) {
	case bool:
		value := GetFlagBool(cmd, flag)
		return value, value == expectValue.(bool)
	case string:
		value := GetFlagString(cmd, flag)
		return value, value == expectValue
	}
	return nil, false
}

func GetFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}

	return s
}

func GetFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.PersistentFlags().GetBool(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}

	return b
}
