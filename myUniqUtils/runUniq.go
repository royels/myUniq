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
		INFILE, err := os.OpenFile(info.Infile, os.O_RDONLY, 0644)
		if err != nil {
			//SetErrorInfo(errorInfo, constants.ErrErrno_M, info.Infile)
			return
		}
		defer INFILE.Close()
	}
	if info.Outfile != "" {
		OUTFILE, err := os.OpenFile(info.Outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
			SetErrorInfo(errorInfo, constants.ErrErrno_M, info.Outfile)
		}
		defer OUTFILE.Close()
	}

	parsedInfo := constants.ParsedInputInfo{}
	retVal := ParseInputStream(INFILE, info, &parsedInfo, errorInfo)
	if retVal != 0 {
		return
	}

	uniqInfo := constants.UniqInfo{}
	retVal = findUniq(&parsedInfo, info, &uniqInfo, errorInfo)
	if retVal != 0 {
		return
	}
	PrintResults(OUTFILE, info, &uniqInfo)
}

//
//void runUniq( const struct argInfo * argInfoPtr,
//struct errorInfo * errorInfoPtr ) {
//
//  FILE* INFILE;
//  FILE* OUTFILE;
//  int i = 0;
//  int j = 0;
//
//
//  // handles infile and outfile errors, otherwise
//  // progresses.
//  INFILE = stdin;
//  OUTFILE = stdout;
//  if(argInfoPtr->inFile != NULL) {
//    INFILE = fopen(argInfoPtr->inFile, FILE_READ_MODE);
//  }
//  if(INFILE == NULL) {
//    SetErrorInfo(errorInfoPtr, ErrErrno_M, argInfoPtr->inFile);
//    return;
//  }j
//  if(argInfoPtr->outFile != NULL) {
//    OUTFILE = fopen(argInfoPtr->outFile, FILE_WRITE_MODE);
//  }
//  if(OUTFILE == NULL) {
//    SetErrorInfo(errorInfoPtr, ErrErrno_M, argInfoPtr->outFile);
//    return;
//  }
//
//
//  struct parsedInputInfo parsedInputInfoPtr;
//
//  // runs parsedInputStream, handles errors.
//
//  int retVal = parseInputStream(INFILE, argInfoPtr,
//  &parsedInputInfoPtr, errorInfoPtr);
//  if(retVal == 1) {
//    for(i = 0; i < parsedInputInfoPtr.numOfEntries; i++) {
//      free(parsedInputInfoPtr.parsedInputPtr[i]);
//    }
//    free(parsedInputInfoPtr.parsedInputPtr);
//    return;
//  }
//
//  // runs finduniq, handles errors
//  struct uniqInfo uniqInfoPtr;
//  retVal = findUniq(&parsedInputInfoPtr, argInfoPtr,
//  &uniqInfoPtr, errorInfoPtr);
//  if(retVal == 1) {
//    for(i = 0; i < parsedInputInfoPtr.numOfEntries; i++) {
//      free(parsedInputInfoPtr.parsedInputPtr[i]);
//    }
//    free(parsedInputInfoPtr.parsedInputPtr);
//    for(j = 0; j < uniqInfoPtr.numOfEntries; j++) {
//      uniq_t temporary = uniqInfoPtr.uniqPtr[j];
//      free(temporary.dups);
//    }
//    free(uniqInfoPtr.uniqPtr);
//    return;
//  }
//
//  // runs printResults, handles errors.
//  printResults(OUTFILE, argInfoPtr, &uniqInfoPtr);
//  for(i = 0; i < parsedInputInfoPtr.numOfEntries; i++) {
//    free(parsedInputInfoPtr.parsedInputPtr[i]);
//  }
//  free(parsedInputInfoPtr.parsedInputPtr);
//
//  for(j = 0; j < uniqInfoPtr.numOfEntries; j++) {
//    uniq_t temporary = uniqInfoPtr.uniqPtr[j];
//    free(temporary.dups);
//  }
//  free(uniqInfoPtr.uniqPtr);
//  (void)fclose(INFILE);
//  (void)fclose(OUTFILE);
//
//}
