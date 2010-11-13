# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=git.78762.de/go/genetics
GOFILES=\
	algorithm.go\
	genome.go\
	initializer.go\
	subject.go\
	evaluator.go\
	selector.go\
	breeder.go\
	mutator.go\

include $(GOROOT)/src/Make.pkg
