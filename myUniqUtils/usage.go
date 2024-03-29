package myuniq

import "fmt"
import (
	"io"
	"pa4/constants"
)

func Usage(w io.Writer, usageMode constants.UsageMode, programName string) {

	if usageMode == constants.UsageShort {
		fmt.Fprintf(w, constants.STR_USAGE_SHORT, programName)
	} else {
		fmt.Fprintf(w, constants.STR_USAGE_LONG, programName)
	}
}
