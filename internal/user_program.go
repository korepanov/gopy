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

type dictionaryValues struct {
	builtinFunctions []string
	delimiters       []string
}

func getDictionary() (d struct {
	dictionaryValues
	getEverything func() []string
}) {

	d.builtinFunctions = []string{"print", "len"}
	d.delimiters = []string{"(", ")", "[", "]", " ", "    "}

	d.getEverything = func() []string {
		dictionaryValues.
	}

	return d
}
