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
	 * Command line flags (most correspond to options, but some set outputMode)
	 */
	FLAG_DUP_ONLY    = 'd'
	FLAG_DUP_ALL     = 'D'
	FLAG_HELP        = 'h'
	FLAG_IGNORE_CASE = 'i'
	FLAG_SORT_OUTPUT = 's'
	FLAG_SORT_INPUT  = 'S'
	FLAG_COUNT       = 'c'
	FLAG_SUMMARY     = 'x'
	FLAG_UNIQUE      = 'u'

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
	options    int
	outputMode OutputMode
	infile     string
	outfile    string
}

/* Used to pass error code and message */
type ErrorInfo struct {
	errorCode int
	errorMsg  string
}

/* Used to keep track of each group of successive, duplicate lines */
type Uniq struct {
	count int
	line  string
	dups  string /* To hold duplicate lines for -D option */
}

/* Typedef this data structure so we can easily modify implementation later. */

/* Stores information about array of uniq_t for whole file */
type UniqInfo struct {
	uniqPtr      Uniq
	numOfEntries int
}

/* Stores information about array of pointers to lines of input file */
type ParsedInputInfo struct {
	parsedInput  string
	numOfEntries int
}

///* Function Prototypes */
//int compareCount(const void *p1, const void *p2);
//
//int compareLine(const void *p1, const void *p2);
//
//
//type pa4Vanilla interface {
//    findUniq(
//        ParsedInputInfo,
//        ArgInfo,
//        UniqInfo,
//        ErrorInfo) int
//    parseArgs(int, ArgInfo,
//        ErrorInfo)
//    parseInputStream(
//            FILE *stream,
//            ArgInfo,
//        struct parsedInputInfo *parsedInputInfoPtr,
//        struct errorInfo *errorInfoPtr
//        );
//            void printErrors(const struct errorInfo *errorInfoPtr, const char *progName);
//            void printResults(
//            FILE *outFile,
//            const struct argInfo *argInfoPtr,
//            const struct uniqInfo *uniqInfoPtr
//        );
//            void runUniq(const struct argInfo *argInfoPtr, struct errorInfo *errorInfoPtr);
//            int sortInputCompare (const void *p1, const void *p2);
//            void setErrorInfo(
//        struct errorInfo *errorInfoPtr,
//            enum errorCode e,
//            const char *errorMsg
//        );
//            int strIgnoreCaseCmp(const char *s1, const char *s2, unsigned int n);
//            void usage(FILE *stream, enum usageMode u, const char *progName);
//        }
