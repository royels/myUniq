package myuniq

import (
	"pa4/constants"
	"fmt"
	"os"
)



func PrintErrors (errorInfoPtr *constants.ErrorInfo, progName string ) {
	if(errorInfoPtr.errorCode == constants.ErrNone) {
		return;
	} else if(errorInfoPtr.errorCode == constants.ErrErrno_M) {
		fmt.Println(errorInfoPtr.errorMsg);
	} else if(errorInfoPtr.errorCode == constants.ErrExtraArgs_M) {
		fmt.Fprint(os.Stderr, constants.STR_ERR_EXTRA_ARGS, errorInfoPtr.errorMsg);
	} else if(errorInfoPtr.errorCode == constants.ErrMutualExcl) {
		fmt.Fprintf(os.Stderr, constants.STR_ERR_EXTRA_ARGS);
	} else {
		fmt.Fprint(os.Stderr, constants.STR_USAGE_SHORT, progName);
	}
}