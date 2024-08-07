package command

const (
	Delimiter lexemeType = iota + 1
	BuiltinFunction
	UserName
	Keyword
	Operator
	Escape
	QuotedString
)

func Dictionary() (d struct {
	dictionaryValues map[lexemeType][]lexeme
	stopList         []lexemeType
	IsStop           func(t lexemeType) bool
	Find             func(lex lexeme) lexemeType
}) {

	d.dictionaryValues = make(map[lexemeType][]lexeme)
	d.dictionaryValues[BuiltinFunction] = []lexeme{"print", "len"}
	d.dictionaryValues[Delimiter] = []lexeme{"(", ")", "[", "]", " ", "    "}
	d.dictionaryValues[Operator] = []lexeme{"+", "-", "*", "**", "/", "//", "%",
		"<", ">", "<=", ">=", "==", "!="}
	d.dictionaryValues[Keyword] = []lexeme{"False", "else", "import", "pass",
		"None", "break", "except", "in", "raise",
		"True", "finally", "is", "return",
		"and", "continue", "for", "try",
		"as", "def", "from", "while",
		"global", "not",
		"elif", "if", "or"}

	d.dictionaryValues[Escape] = []lexeme{" ", "\n", "\t", "\r"}
	d.stopList = []lexemeType{Delimiter, Operator}

	d.Find = func(lex lexeme) lexemeType {

		for key, vals := range d.dictionaryValues {
			for _, val := range vals {
				if val == lex {
					return key
				}
			}
		}

		if len(lex) > 1 && lex[0] == '"' && lex[len(lex)-1] == '"' {
			return QuotedString
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

	return
}
