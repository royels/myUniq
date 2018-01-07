/*
 * Filename: printErrors.c
 * Author: Rohan Yelsangikar
 * Userid: cs30xkr
 * Description: C routine to print all erros
 * Date: 3/05/2015
 * Sources of Help: CSE30 website, man pages
 */

//#include <string.h>
//#include <stdlib.h>
//#include "pa4Strings.h"
//#include "pa4.h"
/*
 * Function name: printErrors()
 * Function prototype:
 * Description: prints all results  made in the duration of this project.
 * Parameters:
 *      arg1: the stream to print to
 *      arg2: the argInfo struct that contains
 *      program argument information
 *      arg3: the uniqInfoPtr that contains information
 *      to be printed
 * Side effects: outputs the results stated by uniqInfo
 * Return Value: None.
 */
package myuniq

import (
	"fmt"
	"os"
	"pa4/constants"
	"strings"
)

type Uniq struct {
	count int
	line  string
	dups  []string /* To hold duplicate lines for -D option */
}

func PrintResults(outFile *os.File, info *constants.ArgInfo, uniqInfo *constants.UniqInfo) {
	regularLineCount := 0
	duplicateLineCount := 0
	duplicateEntryCount := 0
	uniqueEntryCount := 0
	uniqueLineCount := 0
	regularEntryCount := 0
	for i := 0; i < uniqInfo.NumOfEntries; i++ {
		temp := uniqInfo.UniqPtr[i]
		switch info.OutputMode {
		case constants.Regular:
			if (info.Options & constants.OPT_COUNT) != 0 {
				fmt.Fprintf(outFile, constants.STR_PRINT_COUNT, temp.Count)
			}
			fmt.Fprintf(outFile, constants.STR_PRINT_LINE, temp.Line)

			if temp.Count == 1 {
				uniqueLineCount++
				uniqueEntryCount++
			} else if temp.Count > 1 {
				uniqueLineCount++
				duplicateEntryCount++
				duplicateLineCount += temp.Count - 1
			}
			break

		case constants.DuplicateOnly:
			if temp.Count > 1 {
				if (info.Options & constants.OPT_COUNT) != 0 {
					fmt.Fprintf(outFile, constants.STR_PRINT_COUNT, temp.Count)
				}
				fmt.Fprintf(outFile, constants.STR_PRINT_LINE, temp.Line)
			}

			if temp.Count == 1 {
				uniqueLineCount++
				uniqueEntryCount++
			} else if temp.Count > 1 {
				uniqueLineCount++
				duplicateEntryCount++
				duplicateLineCount += temp.Count - 1
			}
			break
		case constants.DuplicateAll:
			if temp.Count > 1 {
				fmt.Fprintf(outFile, temp.Line)
				fmt.Fprintf(outFile, strings.Join(temp.Dups, ","))
			}
			if temp.Count == 1 {
				uniqueEntryCount++
				uniqueLineCount++
			} else if temp.Count > 1 {
				uniqueLineCount++
				duplicateEntryCount++
				duplicateLineCount += temp.Count - 1
			}
			break

		case constants.Unique:
			if temp.Count == 1 {
				if (info.Options & constants.OPT_COUNT) != 0 {
					fmt.Fprintf(outFile,
						constants.STR_PRINT_COUNT, temp.Count)
				}
				fmt.Fprintf(outFile, constants.STR_PRINT_LINE, temp.Line)
			}
			if temp.Count == 1 {
				uniqueLineCount++
				uniqueEntryCount++
			} else if temp.Count > 1 {
				uniqueLineCount++
				duplicateEntryCount++
				duplicateLineCount += temp.Count - 1
			}
			break
		default:
			break
		}
	}

	// If a summary flag has been set
	if (info.Options & constants.OPT_SUMMARY) != 0 {
		fmt.Fprintf(outFile, constants.STR_PRINT_SUMMARY,
			duplicateEntryCount,
			uniqueEntryCount,
			regularEntryCount+duplicateEntryCount+uniqueEntryCount,
			duplicateLineCount,
			uniqueLineCount,
			regularLineCount+uniqueLineCount+duplicateLineCount)

	}

}
