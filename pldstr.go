package main

type PLD_strtok struct {
	start     *string
	delimeter string
}

type PLD_strreplace struct {
	source      *string
	searchfor   *string
	replacewith *string

	preexist  *string
	postexist *string

	replacenumber int

	insensitive int
}
