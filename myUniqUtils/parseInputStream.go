/*
 * Filename: parseInputStream.c
 * Author: Rohan Yelsangikar
 * Userid: cs30xkr
 * Description: C routine to get information out of the user
 * Date: 3/04/2015
 * Sources of Help: CSE30 website, man pages
 */

/*
 * Function name: parseInputStream
 * Function prototype:
 *  int parseInputStream( FILE * stream,
 *  const struct argInfo * argInfoPtr,
 *  struct parsedInputInfo * parsedInputInfoPtr,
 *  struct errorInfo * errorInfoPtr )
 * Description:  gets values out of the user
 * Parameters:
 *      arg1: stream to get values from
 *      arg2: the struct of information about the arguments.
 *      arg3: the struct parsedInputPtr to be populated
 *      arg4: the struct errorInfo to be populated if error.
 * Error conditions: if realloc/calloc fails.
 * Return Value: zero if success, nonzero otherwise
 */
package myuniq

import (
	"io/ioutil"
	"os"
	"pa4/constants"
	"sort"
	"strings"
)

func ParseInputStream(file *os.File, info *constants.ArgInfo, inputInfo *constants.ParsedInputInfo, errorInfo *constants.ErrorInfo) int {
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return 1
	}
	lines := strings.Split(string(content), "\n")

	if info.Options&constants.OPT_SORT_INPUT != 0 {
		sort.Strings(lines)
	}
	inputInfo.ParsedInput = lines
	inputInfo.NumOfEntries = len(lines)
	return 0
}