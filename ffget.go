package main

/* DEFINES */

var FFGET_VERSION	"1.0.0.8"
var FFGET_LASTUPDATED "200811061423"

var FFGET_DNORMAL   ((FFGET_debug >= FFGET_DEBUG_NORMAL  ))
var FFGET_DPEDANTIC ((FFGET_debug >= FFGET_DEBUG_PEDANTIC))

var FFGET_MAX_LINE_LEN 1024
var FFGET_BUFFER_MAX 8192
var FFGET_BUFFER_PADDING 1

var FFGET_DEBUG_NORMAL 1
var FFGET_DEBUG_PEDANTIC 10

var FFGET_LINEBREAK_NONE 0
var FFGET_LINEBREAK_LF 1
var FFGET_LINEBREAK_CR 2

struct _FFGET_FILE
{
	FILE *f;
	char buffer[FFGET_BUFFER_MAX+4];
	char *startpoint;
	char *endpoint;
	char *buffer_end;
	size_t last_block_read_from;
	int FILEEND;
	int FFEOF;
	char c;
	unsigned long int bytes;
	unsigned long int linecount;
	int ungetcset;
	int trueblank;
	char lastchar;
	int linebreak;
	char lastbreak[10];


};

typedef struct _FFGET_FILE FFGET_FILE;

// Special Flag to indicate a Double CR Line.
extern int FFGET_doubleCR;
extern int FFGET_SDL_MODE;  // Single Char Delimeter
extern char SDL_MODE_DELIMITS[];
extern char NORM_MODE_DELIMITS[];
extern char *DELIMITERS;

