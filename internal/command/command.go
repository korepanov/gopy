package command

type lexemeType int
type lexeme string

type Token struct {
	Lex lexeme
	T   lexemeType
}

type Command struct {
	Input    string
	subinput string // to copy Input and work with it inside package
	Tokens   []Token
}
