include  $(GOROOT)/src/Make.inc

TARG=ocl

CGOFILES=\
	ocl.go\

CGO_LDFLAGS=-lOpenCL

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
