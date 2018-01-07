/*
 * Filename: pa4.h
 * author: TODO
 * userid: TODO
 * Description: Header file for pa4 assignment
 */

/*
 * Options bitmasks
 */
package constants

const (
	OPT_COUNT       = 0x01
	OPT_IGNORE_CASE = 0x02
	OPT_SORT_OUTPUT = 0x04
	OPT_SORT_INPUT  = 0x08
	OPT_HELP        = 0x10
	OPT_SUMMARY     = 0x20



	/*
	 * Arg string for getopt
	 */
	ARG_STRING_NON_EC = "cdDhisSux"

	/*
	 * File open modes.
	 */
	FILE_READ_MODE  = "rb"
	FILE_WRITE_MODE = "wb"
	ERROR_MSG_SIZE  = 1024
)

/*
 * Enums
 */
/* Which version of the usage to print */

type UsageMode int

const (
	UsageShort UsageMode = iota
	UsageLong
)

/* What error was encountered */
type ErrorCode int

const (
	ErrNone ErrorCode = iota
	ErrMutualExcl
	ErrInvFlag
	ErrErrno_M
	ErrExtraArgs_M
)

/* How to print output */

type OutputMode int

const (
	DuplicateAll  OutputMode = iota /* -D flag: like -d but print all duplicate lines. */
	DuplicateOnly                   /* -d flag: only print duplicate lines.        */
	Unique                          /* -u flag: only print unique lines.           */
	Regular                         /* default mode: print 1st copy of each line   */
)

/*
 * Structs
 */
/* Stores information parsed from command line */
type ArgInfo struct {
	Options    int
	OutputMode OutputMode
	Infile     string
	Outfile    string
}

/* Used to pass error code and message */
type ErrorInfo struct {
	ErrorCode ErrorCode
	ErrorMsg  string
}


/* Used to keep track of each group of successive, duplicate lines */
type Uniq struct {
	Count int
	Line  string
	Dups  []string /* To hold duplicate lines for -D option */
}

/* Typedef this data structure so we can easily modify implementation later. */

/* Stores information about array of uniq_t for whole file */
type UniqInfo struct {
	UniqPtr      []Uniq
	NumOfEntries int
}

/* Stores information about array of pointers to lines of input file */
type ParsedInputInfo struct {
	ParsedInput  []string
	NumOfEntries int
}