package main

var LIBAM_VERSION = "200811161138"

var (
	AM_RETURN_SIGNED_EMAIL      = 10
	AM_RETURN_B64_ENCODED_EMAIL = 12
)

var (
	AM_DISCLAIMER_TYPE_NONE     = 0
	AM_DISCLAIMER_TYPE_FILENAME = 1
	AM_DISCLAIMER_TYPE_TEXT     = 2
)

var (
	AM_PRETEXT_TYPE_NONE     = 0
	AM_PRETEXT_TYPE_FILENAME = 1
	AM_PRETEXT_TYPE_TEXT     = 2
)

var (
	AM_HEADER_ADJUST_MODE_NONE    = 0
	AM_HEADER_ADJUST_MODE_PREFIX  = 1
	AM_HEADER_ADJUST_MODE_SUFFIX  = 2
	AM_HEADER_ADJUST_MODE_REPLACE = 4
)

var (
	AM_NULLIFY_MATCH_MODE_NONE         = 0
	AM_NULLIFY_MATCH_MODE_FILENAME     = 1
	AM_NULLIFY_MATCH_MODE_CONTENT_TYPE = 2
)

type AM_disclaimer_details struct {
	// Header details

	content_type     int
	content_encoding int
	boundary_found   int
	boundary         [1024]byte

	//

	isb64         int
	ishtml        int
	isfile        int
	text_inserted int
	html_inserted int
	b64_inserted  int

	//

	disclaimer_text_plain *string
	disclaimer_text_HTML  *string
	disclaimer_text_b64   *string

	/** Positional definitions for the HTML and text disclaimers **/
	textpos [1024]byte
	htmlpos [1024]byte
}

var (
	AM_HEADERBUFFER_MAX       = 100
	AM_HEADERBUFFER_ITEM_SIZE = 1024
)

type AM_globals struct {
	debug              int // Low level debugging
	verbose            int /* do we talk as we walk */
	paranoid           int /* set paranoid to yes! */
	HTML_too           int /* Add footer to the HTML email too */
	force_for_bad_html int /** Force insertion of HTML disclaimer even when we can't find the end **/
	force_into_b64     int /** Force headers into Base64 encoded bodies **/
	multipart_insert   int /* Should we insert into emails which are embedded into another */
	nullify_all        int /* Remove ALL filename'd attachments */
	alter_signed       int /* Do we alter signed emails ? */
	header_long_search int /* do we search through email bodies for more headers, like qmail bounces */

	ldelimeter [3]byte

	disclaimer_plain      *string
	disclaimer_plain_type int

	disclaimer_HTML      *string
	disclaimer_HTML_type int

	disclaimer_b64      *string
	disclaimer_b64_type int

	pretext_plain      *string
	pretext_plain_type int
	pretext_HTML       *string
	pretext_HTML_type  int

	pretext_insert int

	headerbuffer [AM_HEADERBUFFER_MAX]byte // 100 lines for the header buffers

	headerbuffermax int
}

var (
	AMSTATUSFLAGS_TEXT_INSERTED    = 1
	AMSTATUSFLAGS_HTML_INSERTED    = 2
	AMSTATUSFLAGS_B64_INSERTED     = 3
	AMSTATUSFLAGS_XHEADER_INSERTED = 4
)

// extern unsigned int altermime_status_flags; // Status flags
