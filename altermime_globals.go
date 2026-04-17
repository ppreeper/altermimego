package main

// Global variables/settings

var (
	ALTERMIMEAPP_default_remove_prefix string = "removed"
	ALTERMIMEAPP_removeall_filename    string = ".*"
)

type ALTERMIMEAPP_globals struct {
	tmpdir                [1024]byte
	input_file            string
	input_is_stdin        int
	disclaimer_file       string
	disclaimer_html_file  string
	disclaimer_b64_file   string
	disclaimer_attachment string
	disclaimer_insert     int
	pretext_file          string
	pretext_html_file     string
	pretext_b64_file      string
	pretext_insert        int
	remove_filename       string
	replace               string
	with                  string
	xheader               string
	alter_header          string
	alter_with            string
	alter_mode            int
	verbose               int
}

func ALTERMIMEAPP_init() *ALTERMIMEAPP_globals {
	return &ALTERMIMEAPP_globals{
		input_file:           "",
		input_is_stdin:       0,
		disclaimer_file:      "",
		disclaimer_html_file: "",
		disclaimer_b64_file:  "",
		disclaimer_insert:    0,
		pretext_file:         "",
		pretext_html_file:    "",
		pretext_b64_file:     "",
		pretext_insert:       0,
		remove_filename:      "",
		replace:              "",
		with:                 "",
		xheader:              "",
		alter_header:         "",
		alter_with:           "",
		alter_mode:           AM_HEADER_ADJUST_MODE_NONE,
		verbose:              0,
	}
}
