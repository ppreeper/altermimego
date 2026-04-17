package main

var SS_STRLEN_MAX = 1024

type SS_node struct {
	data        *string
	data_length size_t
	SS_node     *next
}

type SS_object struct {
	debug         int
	verbose       int
	count         int
	detect_limit  int
	SS_node       *stringstack
	datastacksafe [SS_STRLEN_MAX]byte
}
