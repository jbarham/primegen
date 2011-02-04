include $(GOROOT)/src/Make.inc

TARG=github.com/jbarham/primegen.go

GOFILES=\
	primegen.go\
	count.go\
	fill.go\
	sieve.go\

include $(GOROOT)/src/Make.pkg
