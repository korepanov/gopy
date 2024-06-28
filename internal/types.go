package internal

type lexemeType int
type lexeme string

type Program struct {
	input []command
}

type command struct {
	input  string
	tokens []token
}

type token struct {
	lex lexeme
	t   lexemeType
}
