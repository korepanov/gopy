package internal

func (c *command) lexicalAnalyze() {
	var lexeme string

	for _, ch := range c.input {
		lexeme += string(ch)

	}
}
