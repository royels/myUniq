package myuniq

import "pa4/constants"
import "flag"
import (
	"fmt"
	"os"
)

func ParseArgs(info *constants.ArgInfo, errorInfo *constants.ErrorInfo) int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, constants.STR_USAGE_LONG)
	}
	optind := 0
	dFlag := 0
	DFlag := 0
	uFlag := 0
	SetErrorInfo(errorInfo, constants.ErrNone, "")
	info.OutputMode = constants.Regular
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

	if *OPT_COUNT_LONG || *OPT_COUNT_SHORT {
		info.Options = info.Options | constants.OPT_COUNT
	}
	if *OPT_IGNORE_CASE_LONG || *OPT_IGNORE_CASE_SHORT {
		info.Options = info.Options | constants.OPT_IGNORE_CASE
	}
	if *OPT_SORT_OUTPUT_LONG || *OPT_SORT_OUTPUT_SHORT {
		info.Options = info.Options | constants.OPT_SORT_OUTPUT
	}
	if *OPT_SORT_INPUT_LONG || *OPT_SORT_INPUT_SHORT {
		info.Options = info.Options | constants.OPT_SORT_INPUT
	}
	if *OPT_SUMMARY_LONG || *OPT_SUMMARY_SHORT {
		info.Options = info.Options | constants.OPT_SUMMARY
	}
	if *OPT_HELP_LONG || *OPT_HELP_SHORT {
		info.Options = info.Options | constants.OPT_HELP
	}
	if *OPT_DUP_ONLY_LONG || *OPT_DUP_ONLY_SHORT {
		info.OutputMode = constants.DuplicateOnly
		dFlag = 1
	}
	if *OPT_COUNT_DUP_ALL_LONG || *OPT_COUNT_DUP_ALL_SHORT {
		info.OutputMode = constants.DuplicateAll
		DFlag = 1
	}
	if *OPT_UNIQUE_LONG || *OPT_UNIQUE_SHORT {
		info.OutputMode = constants.Unique
		uFlag = 1
	}

	if optind < len(flag.Args()) {
		info.Infile = flag.Arg(optind)
		optind++
		if optind < len(flag.Args()) {
			info.Outfile = flag.Arg(optind)
			optind++
		}
	}
	//fmt.Println("The outfile is: " + info.Outfile)
	//fmt.Println("The infile is: " + info.Infile)
	extraFlag := 0
	if optind < len(flag.Args()) {
		extraFlag = 1
	}
	if (uFlag == 1 && dFlag == 1) ||
		(uFlag == 1 && DFlag == 1) ||
		(dFlag == 1 && DFlag == 1) {
		SetErrorInfo(errorInfo, constants.ErrMutualExcl, "")
		return 1
	} else if extraFlag == 1 {
		SetErrorInfo(errorInfo, constants.ErrExtraArgs_M, os.Args[optind])
		return 1

	} else if (info.Options & constants.OPT_HELP) != 0 {
		fmt.Fprintf(os.Stdout, constants.STR_USAGE_LONG)
		return 1
	}
	return 0
}
