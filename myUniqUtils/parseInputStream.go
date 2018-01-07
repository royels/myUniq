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
	"fmt"
	"io/ioutil"
	"os"
	"pa4/constants"
	"sort"
	"strings"
)

func ParseInputStream(file *os.File, info *constants.ArgInfo, inputInfo *constants.ParsedInputInfo, errorInfo *constants.ErrorInfo) int {
	//fmt.Println("enterign readfile")
	content, err := ioutil.ReadFile(file.Name())
	//fmt.Println("leaving readfile")
	if err != nil {
		fmt.Println("There was an error")
		return 1
	}
	lines := strings.Split(string(content), "\n")
	//fmt.Print("The lines are ")
	//fmt.Println(lines)
	if (info.Options & constants.OPT_SORT_INPUT) != 0 {
		sort.Strings(lines)
	}
	//fmt.Println("Reached inputInfo parsing")
	inputInfo.ParsedInput = lines
	//fmt.Print("The inputinfo is " )
	//fmt.Println(inputInfo.ParsedInput)
	inputInfo.NumOfEntries = len(lines)
	return 0
}
