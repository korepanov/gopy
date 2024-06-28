package internal

const (
	delimiter lexemeType = iota + 1
	builtinFunction
	userName
	keyword
	operator
	escape
)

func dictionary() (d struct {
	dictionaryValues map[lexemeType][]lexeme
	stopList         []lexemeType
	IsStop           func(t lexemeType) bool
	Find             func(lex lexeme) lexemeType
}) {
	d.dictionaryValues = make(map[lexemeType][]lexeme)
	d.dictionaryValues[builtinFunction] = []lexeme{"print", "len"}
	d.dictionaryValues[delimiter] = []lexeme{"(", ")", "[", "]", " ", "    "}
	d.dictionaryValues[operator] = []lexeme{"+", "-", "*", "**", "/", "//", "%",
		"<", ">", "<=", ">=", "==", "!="}
	d.dictionaryValues[keyword] = []lexeme{"False", "else", "import", "pass",
		"None", "break", "except", "in", "raise",
		"True", "finally", "is", "return",
		"and", "continue", "for", "try",
		"as", "def", "from", "while",
		"global", "not",
		"elif", "if", "or"}

	d.dictionaryValues[escape] = []lexeme{" ", "\n", "\t", "\r"}
	d.stopList = []lexemeType{delimiter, operator}

	d.Find = func(lex lexeme) lexemeType {
		for key, vals := range d.dictionaryValues {
			for _, val := range vals {
				if val == lex {
					return key
				}
			}
		}

		return 0
	}

	d.IsStop = func(t lexemeType) bool {
		for _, val := range d.stopList {
			if val == t {
				return true
			}
		}
		return false
	}

	return d
}
