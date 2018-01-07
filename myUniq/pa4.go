package main

import (
	"os"
	"pa4/constants"
	"pa4/myUniqUtils"
)

func main() {
	info := constants.ArgInfo{}
	errorInfo := constants.ErrorInfo{}
	myuniq.ParseArgs(&info, &errorInfo)
	myuniq.RunUniq(&info, &errorInfo)
	myuniq.PrintErrors(&errorInfo, os.Args[0])
}
