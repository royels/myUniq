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
  "io"
  "fmt"
  "pa4/constants"
)

func ParseInputStream (w io.Writer, info * constants.ArgInfo, inputInfo* constants.ParsedInputInfo, errorInfo * constants.ErrorInfo) int {


}


//int parseInputStream( FILE * stream,
//const struct argInfo * argInfoPtr,
//struct parsedInputInfo * parsedInputInfoPtr, struct errorInfo * errorInfoPtr ){
//
//  char **storage;
//  char **storageTemp = malloc(sizeof(char*));
//  if(storageTemp == NULL) {
//    setErrorInfo(errorInfoPtr, ErrErrno_M, STR_ERR_PARSE_INPUT);
//    return 1;
//  }
//
//  // copy over pointers with every realloc.
//  char temp[BUFSIZ];
//  int i = 0;
//  while(fgets(temp, BUFSIZ, stream) != NULL) {
//    char *charTemp  = calloc(strlen(temp) + 1, sizeof(char));
//  if(charTemp == NULL) {
//    setErrorInfo(errorInfoPtr, ErrErrno_M, STR_ERR_PARSE_INPUT);
//    return 1;
//  }
//  (void)strncpy(charTemp, temp, strlen(temp)+1);
//  storage = (char **) realloc(storageTemp,
//  (i+1) * sizeof(char*));
//  if(storage == NULL) {
//    int j = 0;
//    for(j = 0; j < i; j++) {
//        free(storageTemp[j]);
//    }
//    free(storage);
//    setErrorInfo(errorInfoPtr, ErrErrno_M, STR_ERR_PARSE_INPUT);
//    return 1;
//  }
//  storage[i] = charTemp;
//  storageTemp = storage;
//  i++;
// }
//
//  // sort if sort_input flag present
//  if((argInfoPtr->options & OPT_SORT_INPUT) != 0) {
//    qsort(storageTemp, i, sizeof(char*), sortInputCompare);
//  }
//
//
// parsedInputInfoPtr->parsedInputPtr = storageTemp;
// parsedInputInfoPtr->numOfEntries = i;
//
//
// return 0;
//}
