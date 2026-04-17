package main

import (
	"flag"
	"fmt"
)

func main() {
	glb := ALTERMIMEAPP_init()
	// ALTERMIMEAPP_init( &glb );

	// LOGGER_set_output_mode(_LOGGER_STDOUT);

	ALTERMIMEAPP_parse_args(glb)
	// ALTERMIMEAPP_parse_args(&glb,argc, argv);

	// if (!glb.input_file) {
	// 	LOGGER_log("Error: No input file specified\n");
	// 	LOGGER_log(ALTERMIMEAPP_USAGE);
	// 	exit(1);
	// }

	// if (((!glb.replace)&&(glb.with))||((glb.replace)&&(!glb.with)))
	// {
	// 	LOGGER_log("Error: Both --replace= and --with= must be set\n");
	// 	exit(1);
	// }

	// if( glb.input_file && !(glb.alter_mode||glb.replace||glb.disclaimer_file||glb.remove_filename||glb.xheader)) {
	// 	LOGGER_log("Error: Must specify an action for the input file.\n");
	// 	LOGGER_log( ALTERMIMEAPP_USAGE);
	// 	exit(1);
	// }

	// if ((strcmp(glb.input_file,"-") == 0) && (/**glb.disclaimer_file||**/glb.replace||glb.xheader)) {
	// 	LOGGER_log("Error: reading/writing from stdin/stdout not implemented for --xheader,--disclaimer, or --replace.\n");
	// 	LOGGER_log(ALTERMIMEAPP_USAGE);
	// 	exit(1);
	// }

	// if ((glb.alter_mode != AM_HEADER_ADJUST_MODE_NONE)&&(glb.alter_with != NULL)&&(glb.alter_header != NULL))
	// {
	// 	AM_alter_header( glb.input_file, glb.alter_header, glb.alter_with, glb.alter_mode );
	// }

	// 	if ((glb.replace)&&(glb.with)) AM_attachment_replace( glb.input_file, glb.replace, glb.with);
	// 	if (glb.disclaimer_file) AM_add_disclaimer( glb.input_file );
	// #ifdef ALTERMIME_PRETEXT
	// 	if (glb.pretext_insert) AM_add_pretext( glb.input_file );
	// #endif

	// 	if (glb.remove_filename) AM_nullify_attachment(glb.input_file, glb.remove_filename);
	// 	if (glb.xheader) AM_insert_Xheader( glb.input_file, glb.xheader);

	// 	AM_done();

	return
}

func ALTERMIMEAPP_parse_args(glb *ALTERMIMEAPP_globals) error {
	// Define flags
	input := flag.String("input", "", "input file name")
	htmltoo := flag.Bool("htmltoo", false, "insert plaintext disclaimer into HTML")
	forceIntoB64 := flag.Bool("force-into-b64", false, "insert disclaimers into BASE64 encoded text")
	forceBadHtml := flag.Bool("force-for-bad-html", false, "force adding HTML disclaimer even with bad formatting")
	multipartInsert := flag.Bool("multipart-insert", false, "enable multipart insert")

	disclaimerFile := flag.String("disclaimer", "", "plaintext disclaimer source file")
	disclaimerHtml := flag.String("disclaimer-html", "", "HTML disclaimer source file")
	disclaimerB64 := flag.String("disclaimer-b64", "", "BASE64 encoded disclaimer source file")

	removeFile := flag.String("remove", "", "remove file name (regex)")
	removeAll := flag.Bool("removeall", false, "remove all attachments")
	alterSigned := flag.Bool("altersigned", false, "force modify signed emails")

	replace := flag.String("replace", "", "filename to replace")
	with := flag.String("with", "", "replace with filename")
	xheader := flag.String("xheader", "", "insert header line")

	alterHeader := flag.String("alter-header", "", "header to alter")
	alterWith := flag.String("alter-with", "", "new header value")
	alterMode := flag.String("alter-mode", "", "header alter mode (prefix|suffix|replace)")

	debug := flag.Bool("debug", false, "enable debug mode")
	noQmailBounce := flag.Bool("no-qmail-bounce", false, "don't search email bodies for headers")
	verbose := flag.Bool("verbose", false, "verbose output")

	// Parse flags
	flag.Parse()

	// Set values in globals struct
	glb.input_file = *input
	if *htmltoo {
		AM_set_HTMLtoo(1)
	}
	if *forceIntoB64 {
		AM_set_force_into_b64(1)
	}
	if *forceBadHtml {
		AM_set_force_for_bad_html(1)
	}
	if *multipartInsert {
		AM_set_multipart_insert(1)
	}

	glb.disclaimer_file = *disclaimerFile
	if *disclaimerFile != "" {
		AM_set_disclaimer_plain(*disclaimerFile, AM_DISCLAIMER_TYPE_FILENAME)
	}

	glb.disclaimer_html_file = *disclaimerHtml
	if *disclaimerHtml != "" {
		AM_set_disclaimer_HTML(*disclaimerHtml, AM_DISCLAIMER_TYPE_FILENAME)
		AM_set_HTMLtoo(1)
	}

	if *disclaimerB64 != "" {
		glb.disclaimer_b64_file = *disclaimerB64
		AM_set_disclaimer_b64(*disclaimerB64, AM_DISCLAIMER_TYPE_FILENAME)
		AM_set_force_into_b64(1)
	}

	if *removeAll {
		glb.remove_filename = ALTERMIMEAPP_removeall_filename
	} else {
		glb.remove_filename = *removeFile
	}

	if *alterSigned {
		AM_set_altersigned(1)
	}

	glb.replace = *replace
	glb.with = *with
	glb.xheader = *xheader

	if *debug {
		AM_set_debug(1)
	}
	if *noQmailBounce {
		AM_set_header_long_search(0)
	}
	if *verbose {
		AM_set_verbose(1)
	}

	if *alterHeader != "" {
		glb.alter_header = *alterHeader
		glb.alter_with = *alterWith
		switch *alterMode {
		case "prefix":
			glb.alter_mode = AM_HEADER_ADJUST_MODE_PREFIX
		case "suffix":
			glb.alter_mode = AM_HEADER_ADJUST_MODE_SUFFIX
		case "replace":
			glb.alter_mode = AM_HEADER_ADJUST_MODE_REPLACE
		default:
			return fmt.Errorf("unknown header alter mode '%s'. Please use either prefix, suffix or replace", *alterMode)
		}
	}

	return nil
}

func getUsage() string {
	return `altermime --input=<input mime pack>   ( --input=- for stdin )
	[--disclaimer=<disclaimer file>]
	[--disclaimer-html=<HTML disclaimer file>]
	[--disclaimer-b64=<BASE64 encoded dislcaimer>]
	[--htmltoo]
	[--pretext=<pretext file>]
	[--pretext-html=<pretext HTML file>]
	[--force-into-b64]
	[--force-for-bad-html]
	[--multipart-insert]
	[--remove=<remove file name (regex)>] (if filename contains a /, matches on mime-type )
	[--removeall]
	[--replace=<filename to replace> --with=<replace with>]
	[--xheader="..."]
	[--alter-header="..." --alter-with="..." --alter-mode=<prefix|suffix|replace>]
	[--altersigned]
	[--no-qmail-bounce]
	[--verbose]
	[--log-stdout]
	[--log-stderr]
	[--log-syslog]
	[--debug]
	[--version]

Option Descriptions:
	--input=, Sets the mailpack file to be the filename supplied,
		if the filename is a single '-' (hyphen) then the mailpack
		is sourced via stdin and outputted via stdout.

	--disclaimer=, Set the plaintext disclaimer source file.
	--disclaimer-html=, Set the HTML disclaimer source file.
	--disclaimer-b64=, Set the BASE64 encoded disclaimer source file (implies --force-into-b64).

	--htmltoo, Sets alterMIME to insert the plaintext disclaimer into
	--force-into-b64, Sets alterMIME to insert disclaimers into BASE64 encoded text segments
	--force-for-bad-html, Force adding of the HTML disclaimer even when HTML is not correctly formatted
		the HTML portion of the email body ( if there is no explicitly
		defined HTML dislcaimer, see --disclaimer-html )

	--remove=, Remove any attachments which match the filename supplied,
		if the filename text contains a forward-slash '/', then the
		matching will occur based on content-type headers rather than
		by filename.
	--removeall, Remove all attachments
	--replace=, Replace attachments matching the given filename. Requires to
		be used with --with.
	--with=, Replace the attachments specified by --replace with the file
		specified.
	--xheader=, Insert a header line as specified into the first set of headers.
	--alter-header="..." --alter-with="..." --alter-mode=(prefix|suffix|replace)
		Alter an existing header in the mailpack.  This function modifies the
		value of the header, as opposed to the header name.
	--altersigned, Force alterMIME to modify 'signed' emails
	--no-qmail-bounce,  Don't search into email bodies for attachment headers
	--verbose, Describe details of the process occurring
	--log-stdout, Send all output messages to stdout
	--log-stderr, Send all output messages to stderr
	--log-syslog, Send all output messages to syslog
	--debug, Provide greater verbosity and debugging information
	--version, display the alterMIME version string
`
}
