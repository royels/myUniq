/*
 * Filename: SetErrorInfo.c
 * Description: Sets members of sturct errorInfo.
 */

package myuniq

import "pa4/constants"

func SetErrorInfo(errorInfo *constants.ErrorInfo, errorCode constants.ErrorCode, errorMsg string) {

	/* Set error code */
	errorInfo.ErrorCode = errorCode

	/* Set message if not null */
	if errorMsg != "" {
		errorInfo.ErrorMsg = errorMsg
	} else {
		errorInfo.ErrorMsg = ""
	}
}
