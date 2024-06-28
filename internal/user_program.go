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
	delimiter lexemeType = iota + 1
	builtinFunction
	userName
	keyword
	operator
	escape
)

func dictionary() (d struct {
	dictionaryValues map[lexemeType][]string
	StopSequences    func() []string
	Find             func(lexeme string) lexemeType
}) {
	d.dictionaryValues[builtinFunction] = []string{"print", "len"}
	d.dictionaryValues[delimiter] = []string{"(", ")", "[", "]", " ", "    "}
	d.dictionaryValues[operator] = []string{"+", "-", "*", "**", "/", "//", "%",
		"<", ">", "<=", ">=", "==", "!="}
	d.dictionaryValues[keyword] = []string{"False", "else", "import", "pass",
		"None", "break", "except", "in", "raise",
		"True", "finally", "is", "return",
		"and", "continue", "for", "try",
		"as", "def", "from", "while",
		"global", "not",
		"elif", "if", "or"}

	d.dictionaryValues[escape] = []string{" ", "\n", "\t", "\r"}

	d.Find = func(lexeme string) lexemeType {
		for key, vals := range d.dictionaryValues {
			for _, val := range vals {
				if val == lexeme {
					return key
				}
			}
		}

		return 0
	}

	d.StopSequences = func() []string {
		var res []string
		for _, vals := range d.dictionaryValues {

		}
	}
	return d
}
