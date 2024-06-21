package internal

type Program struct {
	InputProgram []Command
}

type Command struct {
	Input  string
	Tokens []Token
}

type Token struct{
	Lexeme string 
	T 
}
type LexemeType string 

const {
	delimiter

}

type Delimiter struct{
	
}