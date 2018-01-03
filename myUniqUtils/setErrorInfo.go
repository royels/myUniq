/*
 * Filename: setErrorInfo.c
 * Description: Sets members of sturct errorInfo.
 */

 package myuniq

import "pa4/constants"

func setErrorInfo(errorInfo * constants.ErrorInfo, errorCode constants.ErrorCode, errorMsg string) {

  /* Set error code */
  errorInfo.errorCode = errorCode;

  /* Set message if not null */
  if(errorMsg != "") {
    errorInfo.errorMsg = errorMsg;
  } else {
    errorInfo.errorMsg = "";
  }
}