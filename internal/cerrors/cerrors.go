package cerrors

import "errors"

var (
	ErrNoToken     = errors.New("no token found")
	ErrLexAnalysis = errors.New("lexical analysis failed")
	ErrRead        = errors.New("reading program failed")
	ErrCompile     = errors.New("compilation failed")
	ErrFlagParse   = errors.New("could not parse flags")
	ErrHelp        = errors.New("help")
)
