package main

var MIMEH_VERSION = "200811151148"

var _CTYPE_RANGE = 99

var (
	_CTYPE_UNSPECIFIED   = -1
	_CTYPE_MESSAGE_START = 1
	_CTYPE_MESSAGE       = 1
	_CTYPE_MESSAGE_END   = 100
)

var (
	_CTYPE_MULTIPART_START       = 100
	_CTYPE_MULTIPART             = 100
	_CTYPE_MULTIPART_MIXED       = 101
	_CTYPE_MULTIPART_APPLEDOUBLE = 102
	_CTYPE_MULTIPART_RELATED     = 103
	_CTYPE_MULTIPART_ALTERNATIVE = 104
	_CTYPE_MULTIPART_REPORT      = 105
	_CTYPE_MULTIPART_SIGNED      = 106
	_CTYPE_MULTIPART_END         = 199
)

var (
	_CTYPE_TEXT_START    = 200
	_CTYPE_TEXT          = 200
	_CTYPE_TEXT_PLAIN    = 201
	_CTYPE_TEXT_UNKNOWN  = 202
	_CTYPE_TEXT_HTML     = 203
	_CTYPE_TEXT_CALENDAR = 204
	_CTYPE_TEXT_END      = 299
)

var (
	_CTYPE_IMAGE_START = 300
	_CTYPE_IMAGE       = 300
	_CTYPE_IMAGE_GIF   = 301
	_CTYPE_IMAGE_JPEG  = 302
	_CTYPE_IMAGE_PNG   = 303
	_CTYPE_IMAGE_END   = 399
)

var (
	_CTYPE_AUDIO_START = 400
	_CTYPE_AUDIO       = 400
	_CTYPE_AUDIO_END   = 499
)

var (
	_CTYPE_OCTECT                = 800
	_CTYPE_RFC822                = 500
	_CTYPE_TNEF                  = 600
	_CTYPE_APPLICATION           = 700
	_CTYPE_APPLICATION_APPLEFILE = 701
	_CTYPE_UNKNOWN               = 0
)

var (
	_CTRANS_ENCODING_UNSPECIFIED = -1
	_CTRANS_ENCODING_B64         = 100
	_CTRANS_ENCODING_7BIT        = 101
	_CTRANS_ENCODING_8BIT        = 102
	_CTRANS_ENCODING_QP          = 103
	_CTRANS_ENCODING_RAW         = 104
	_CTRANS_ENCODING_BINARY      = 105
	_CTRANS_ENCODING_UUENCODE    = 106
	_CTRANS_ENCODING_UNKNOWN     = 0
)

var (
	_CDISPOSITION_UNSPECIFIED = -1
	_CDISPOSITION_INLINE      = 100
	_CDISPOSITION_ATTACHMENT  = 200
	_CDISPOSITION_FORMDATA    = 300
	_CDISPOSITION_UNKNOWN     = 0
)

var _MIMEH_FOUND_FROM = 100

var (
	_MIMEH_STRLEN_MAX                    = 1023
	_MIMEH_FILENAMELEN_MAX               = 128
	_MIMEH_CONTENT_TYPE_MAX              = 128
	_MIMEH_SUBJECTLEN_MAX                = 128
	_MIMEH_CONTENT_DESCRIPTION_MAX       = 128
	_MIMEH_CONTENT_TRANSFER_ENCODING_MAX = 256
	_MIMEH_CONTENT_DISPOSITION_MAX       = 256
	_MIMEH_DEBUG_NORMAL                  = 1
	_MIMEH_DEBUG_PEDANTIC                = 10
	_MIMEH_DEFECT_ARRAY_SIZE             = 100
)

// Errors to throw back
var MIMEH_ERROR_DISK_FULL = 128

// Defects
var (
	MIMEH_DEFECT_MULTIPLE_QUOTES            = 1
	MIMEH_DEFECT_UNBALANCED_QUOTES          = 2
	MIMEH_DEFECT_MULTIPLE_EQUALS_SEPARATORS = 3
	MIMEH_DEFECT_MULTIPLE_COLON_SEPARATORS  = 4
	MIMEH_DEFECT_MULTIPLE_BOUNDARIES        = 5
	MIMEH_DEFECT_UNBALANCED_BOUNDARY_QUOTE  = 6
	MIMEH_DEFECT_MULTIPLE_FIELD_OCCURANCE   = 7
	MIMEH_DEFECT_MISSING_SEPARATORS         = 8
	MIMEH_DEFECT_MULTIPLE_NAMES             = 9
	MIMEH_DEFECT_MULTIPLE_FILENAMES         = 10
)

type MIMEH_header_info struct {
	scratch                    [_MIMEH_STRLEN_MAX + 1]byte
	content_type               int
	content_type_string        [_MIMEH_CONTENT_TYPE_MAX + 1]byte
	content_description_string [_MIMEH_CONTENT_DESCRIPTION_MAX + 1]byte
	boundary                   [_MIMEH_STRLEN_MAX + 1]byte
	boundary_located           int
	subject                    [_MIMEH_SUBJECTLEN_MAX + 1]byte
	filename                   [_MIMEH_FILENAMELEN_MAX + 1]byte
	name                       [_MIMEH_STRLEN_MAX + 1]byte

	//** 20041217-1601:PLD: New header fields to keep **/
	from      [_MIMEH_STRLEN_MAX + 1]byte
	date      [_MIMEH_STRLEN_MAX + 1]byte
	to        [_MIMEH_STRLEN_MAX + 1]byte
	messageid [_MIMEH_STRLEN_MAX + 1]byte
	received  [_MIMEH_STRLEN_MAX + 1]byte
	//** end of new fields **/

	// Store multiple filenames
	SS_object ss_filenames
	// Store multiple names
	SS_object ss_names

	content_transfer_encoding        int
	content_transfer_encoding_string [_MIMEH_CONTENT_TRANSFER_ENCODING_MAX + 1]byte
	content_disposition              int
	content_disposition_string       [_MIMEH_CONTENT_DISPOSITION_MAX + 1]byte
	charset                          int
	format                           int
	file_has_uuencode                int
	uudec_name                       [_MIMEH_FILENAMELEN_MAX + 1]byte // UUDecode name. This is a post-decode information field.
	current_recursion_level          int

	// Malformed email reporting
	defects             [_MIMEH_DEFECT_ARRAY_SIZE]int
	header_defect_count int

	// Special Exception flags
	x_mac int // Set if the content type contains x-mac-* entries, which means a filename may contain /'s

	// Header sanity level - indicates if any of the headers we apparently read are good
	sanity int

	/** 20051117-0932:PLD: Will be non-zero if email is MIME **/
	is_mime int

	delimeter  [3]byte
	crlf_count int // 200811151149:PLD: Tally's the number of CRLF lines
	crcr_count int // 200811151149:PLD: Tally's the number of CRLF lines
	lf_count   int // 200811151149:PLD: Tally's the number of  LF only lines
}

type MIMEH_header_node struct {
	MIMEH_header_info *header_list
	MIMEH_header_node *next
}

type MIMEH_email_info struct {
	mailpack_name     [1024]byte
	MIMEH_header_node *headers
}
