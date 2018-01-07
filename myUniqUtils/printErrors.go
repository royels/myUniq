package myuniq

import (
	"fmt"
	"os"
	"pa4/constants"
)

func PrintErrors(errorInfoPtr *constants.ErrorInfo, progName string) {
	if errorInfoPtr.ErrorCode == constants.ErrNone {
		return
	} else if errorInfoPtr.ErrorCode == constants.ErrErrno_M {
		fmt.Println(errorInfoPtr.ErrorMsg)
	} else if errorInfoPtr.ErrorCode == constants.ErrExtraArgs_M {
		fmt.Fprint(os.Stderr, constants.STR_ERR_EXTRA_ARGS, errorInfoPtr.ErrorMsg)
	} else if errorInfoPtr.ErrorCode == constants.ErrMutualExcl {
		fmt.Fprintf(os.Stderr, constants.STR_ERR_EXTRA_ARGS)
	} else {
		fmt.Fprint(os.Stderr, constants.STR_USAGE_SHORT, progName)
	}
}
