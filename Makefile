# Wedit Makefile for project altermime
#CC=cc
#CC=ccmalloc gcc
#CFLAGS=-Wall -g
#CFLAGS=-Wall -ggdb

# Optional builds
#	ALTERMIME_PRETEXT - Allows prefixing of the email body with a file, sort of the
#								opposite of a disclaimer.
# ALTERMIME_OPTIONS=-DALTERMIME_PRETEXT
#ALTERMIME_OPTIONS=


all: altermime

altermime:
	$go build -o altermime

# Build Install
install: altermime
	# strip altermime
	# cp altermime /usr/local/bin
	# chmod a+rx /usr/local/bin/altermime

# uninstall:
	# rm -f /usr/local/bin/altermime

# clean:
# 	rm -f *.o altermime
