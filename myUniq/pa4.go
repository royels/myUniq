package main

import "fmt"
import (
	"pa4/myUniqUtils"
	"pa4/constants"
	"os"
	"go/constant"
)

func main() {
	argInfo := new(constants.ArgInfo)
	errorInfo := new(constants.ErrorInfo)
	myuniq.ParseArgs(argInfo, errorInfo)
	if(errorInfo.errorCode == constants.ErrNone) {
		//runUniq
	}
	if(errorInfo.errorCode != constants.ErrNone) {
		myuniq.PrintErrors(errorInfo, os.Args[0])
	}
}
