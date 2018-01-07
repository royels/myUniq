package main

import (
	"os"
	"pa4/constants"
	"pa4/myUniqUtils"
)

func main() {
	info := constants.ArgInfo{}
	errorInfo := constants.ErrorInfo{}
	retVal := myuniq.ParseArgs(&info, &errorInfo)
	if(retVal == 1) {
		return
	}
	myuniq.RunUniq(&info, &errorInfo)
	myuniq.PrintErrors(&errorInfo, os.Args[0])
}
