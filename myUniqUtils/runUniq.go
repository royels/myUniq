/*
 * Filename: runUniq.c
 * Author: Rohan Yelsangikar
 * Userid: cs30xkr
 * Description: THe main runner of this asssignment.
 * Date: 3/06/2015
 * Sources of Help: CSE30 website, man pages
 */

/*
 * Function name: runUniq()
 * Function prototype:
 * void runUniq( const struct argInfo * argInfoPtr,
 * struct errorInfo * errorInfoPtr )
 * Description: main delegater for running myUniq
 * Parameters:
 *      arg1: argInfo with information about program args
 *      arg2: struct errorInfo to be populated if error.
 * Error conditions: if fopen, parseInputSTream, or findUniq
 * fails, then an errorInfo is popualted.
 * Return Value: None
 */
package myuniq

import (
	"os"
	"pa4/constants"
)

func RunUniq(info *constants.ArgInfo, errorInfo *constants.ErrorInfo) {

	INFILE := os.Stdin
	OUTFILE := os.Stdout

	if info.Infile != "" {
		val, err := os.OpenFile(info.Infile, os.O_RDONLY, 0644)
		if err != nil {
			//SetErrorInfo(errorInfo, constants.ErrErrno_M, info.Infile)
			return
		} else {
			INFILE = val
		}

	}
	defer INFILE.Close()

	if info.Outfile != "" {
		val, err := os.OpenFile(info.Outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
			SetErrorInfo(errorInfo, constants.ErrErrno_M, info.Outfile)
		} else {
			OUTFILE = val
		}
	}
	defer OUTFILE.Close()

	//fmt.Println("Reached ParseInputStream")
	parsedInfo := constants.ParsedInputInfo{}
	retVal := ParseInputStream(INFILE, info, &parsedInfo, errorInfo)
	//fmt.Println(retVal)
	if retVal != 0 {
		return
	}
	//fmt.Println("Reached findUniq")
	uniqInfo := constants.UniqInfo{}
	//fmt.Print("parsedInfo.ParsedInput is ")
	//fmt.Println(parsedInfo.ParsedInput)
	retVal = findUniq(&parsedInfo, info, &uniqInfo, errorInfo)
	if retVal != 0 {
		return
	}
	//fmt.Println("Reached PrintResutls")
	PrintResults(OUTFILE, info, &uniqInfo)
}
