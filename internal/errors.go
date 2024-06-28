package internal

import "errors"

var (
	ErrNoToken     = errors.New("no token found")
	ErrLexAnalysis = errors.New("lexical analysis failed")
	ErrRead        = errors.New("reading program failed")
	ErrCompile     = errors.New("compiling failed")
)
