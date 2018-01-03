
package myuniq

import "pa4/constants"
import "flag"
import (
	"os"
	"fmt"
)






func ParseArgs(info * constants.ArgInfo, errorInfo * constants.ErrorInfo ) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, constants.STR_USAGE_LONG)
	}
	optind := 0
	 dFlag := 0;
	 DFlag := 0;
	 uFlag := 0;
	 setErrorInfo(errorInfo, constants.ErrNone, "")
	info.outputMode = constants.Regular
	OPT_COUNT_LONG := flag.Bool(constants.STR_LONG_OPT_COUNT, false, "should not come up")
	OPT_IGNORE_CASE_LONG := flag.Bool(constants.STR_LONG_OPT_IGNORE_CASE, false, "should not come up")
	OPT_SORT_OUTPUT_LONG := flag.Bool(constants.STR_LONG_OPT_SORT_OUTPUT, false, "should not come up")
	OPT_SORT_INPUT_LONG := flag.Bool(constants.STR_LONG_OPT_SORT_INPUT, false, "should not come up")
	OPT_SUMMARY_LONG := flag.Bool(constants.STR_LONG_OPT_SUMMARY, false, "should not come up")
	OPT_HELP_LONG := flag.Bool(constants.STR_LONG_OPT_HELP, false, "should not come up")
	OPT_DUP_ONLY_LONG := flag.Bool(constants.STR_LONG_OPT_DUP_ONLY, false, "should not come up")
	OPT_COUNT_DUP_ALL_LONG := flag.Bool(constants.STR_LONG_OPT_DUP_ALL, false, "should not come up")
	OPT_UNIQUE_LONG := flag.Bool(constants.STR_LONG_OPT_UNIQUE, false, "should not come up")

	OPT_COUNT_SHORT := flag.Bool(constants.FLAG_COUNT, false, "should not come up")
	OPT_IGNORE_CASE_SHORT := flag.Bool(constants.FLAG_IGNORE_CASE, false, "should not come up")
	OPT_SORT_OUTPUT_SHORT := flag.Bool(constants.FLAG_SORT_OUTPUT, false, "should not come up")
	OPT_SORT_INPUT_SHORT := flag.Bool(constants.FLAG_SORT_INPUT, false, "should not come up")
	OPT_SUMMARY_SHORT := flag.Bool(constants.FLAG_SUMMARY, false, "should not come up")
	OPT_HELP_SHORT := flag.Bool(constants.FLAG_HELP, false, "should not come up")
	OPT_DUP_ONLY_SHORT := flag.Bool(constants.FLAG_DUP_ONLY, false, "should not come up")
	OPT_COUNT_DUP_ALL_SHORT := flag.Bool(constants.FLAG_DUP_ALL, false, "should not come up")
	OPT_UNIQUE_SHORT := flag.Bool(constants.FLAG_UNIQUE, false, "should not come up")
	flag.Parse()

	if(*OPT_COUNT_LONG || *OPT_COUNT_SHORT) {
		info.options = info.options | constants.OPT_COUNT
	}
	if (*OPT_IGNORE_CASE_LONG || *OPT_IGNORE_CASE_SHORT) {
		info.options = info.options | constants.OPT_IGNORE_CASE
	}
	if(*OPT_SORT_OUTPUT_LONG || *OPT_SORT_OUTPUT_SHORT) {
		info.options = info.options | constants.OPT_SORT_OUTPUT
	}
	if(*OPT_SORT_INPUT_LONG || *OPT_SORT_INPUT_SHORT) {
		info.options = info.options | constants.OPT_SORT_INPUT
	}
	if(*OPT_SUMMARY_LONG || *OPT_SUMMARY_SHORT) {
		info.options = info.options | constants.OPT_SUMMARY
	}
	if(*OPT_HELP_LONG || *OPT_HELP_SHORT) {
		info.options = info.options | constants.OPT_HELP
	}
	if(*OPT_DUP_ONLY_LONG || *OPT_DUP_ONLY_SHORT) {
		info.outputMode = constants.DuplicateOnly
		dFlag = 1
	}
	if(*OPT_COUNT_DUP_ALL_LONG || *OPT_COUNT_DUP_ALL_SHORT) {
		info.outputMode = constants.DuplicateAll
		DFlag = 1
	}
	if(*OPT_UNIQUE_LONG || *OPT_UNIQUE_SHORT) {
		info.outputMode = constants.Unique
		uFlag = 1
	}
	if(optind < len(os.Args)) {
		info.inFile = os.Args[optind];
		optind++;
		if(optind < len(os.Args)) {
			info.outFile = os.Args[optind];
			optind++;
		}
	}
	extraFlag := 0
	if(optind < len(os.Args)) {
		extraFlag = 1;
	}
	if((uFlag == 1 &&  dFlag == 1) ||
		(uFlag == 1 && DFlag == 1) ||
		(dFlag == 1 && DFlag == 1)) {
		setErrorInfo(errorInfo, constants.ErrMutualExcl, "");
	} else if(extraFlag == 1) {
		setErrorInfo(errorInfo, constants.ErrExtraArgs_M, os.Args[optind]);
	}
}