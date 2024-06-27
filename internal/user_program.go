package internal

type Program struct {
	input []command
}

type command struct {
	input  string
	tokens []token
}

type token struct {
	lexeme string
	t      lexemeType
}

type lexemeType int

const (
	delimiter lexemeType = iota
	builtinFunction
	userName
	keyword
)
