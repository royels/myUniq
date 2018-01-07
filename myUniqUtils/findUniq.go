package myuniq

import "pa4/constants"

/*
 * Filename: findUniq.c
 * Author: Rohan Yelsangikar
 * Userid: cs30xkr
 * Description: populates uniqInfoPtr, with values that are
 * complete with a count, dups, and line member.
 * Date: 3/08/2015
 * Sources of Help: CSE30 website, man pages
 */

/*
 * Function name: findUniq()
 * Function prototype: int findUniq(
 * const struct parsedInputInfo * parsedInputInfoPtr,
 * const struct argInfo * argInfoPtr,
 * struct uniqInfo * uniqInfoPtr, struct errorInfo * errorInfoPtr )
 * Description: simulates an anagram creator out of a source file
 * Parameters:
 *    arg1: parsedInputInfoPtr: the parsed information
 *    arg2: argInfoPtr: argumentInfo to determine course of program
 *    arg3: uniqInfoPtr: struct to be populated iwth information
 *    arg4: errorInfoPtr: struct to be populated if error.
 * Error conditions: if realloc returns NULL
 * Return Value: 0 if success, nonzero otherwise.
 */

func findUniq(info *constants.ParsedInputInfo, argInfo *constants.ArgInfo, uniqInfo *constants.UniqInfo, errorInfo *constants.ErrorInfo) int {
	i := 0
	uniq_string_list := []string{}
	uniq_obj_list := []constants.Uniq{}
	uniq_vals := make(map[string]constants.Uniq)
	for j := 0; j < info.NumOfEntries; j++ {
		val, ok := uniq_vals[info.ParsedInput[i]]
		if !ok {
			uniq := constants.Uniq{Count: 1, Line: info.ParsedInput[i], Dups: []string{}}
			uniq_vals[info.ParsedInput[i]] = uniq
			uniq_string_list = append(uniq_string_list, info.ParsedInput[i])
		} else {
			val.Count = val.Count + 1
			val.Dups = append(val.Dups, info.ParsedInput[i])
		}
	}
	for k := 0; k < len(uniq_string_list); k++ {
		uniq_obj_list = append(uniq_obj_list, uniq_vals[uniq_string_list[k]])
	}
	uniqInfo.UniqPtr = uniq_obj_list // should be the ordered list of unique
	uniqInfo.NumOfEntries = len(uniq_obj_list)
	return 0

}
