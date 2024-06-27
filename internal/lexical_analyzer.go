package internal

func (c *command) lexicalAnalyze() {
	var token1 token
	var token2 token

	token1.lexeme = "("
	token1.t = delimiter

	token2.lexeme = "print"
	token2.t = builtinFunction

	c.tokens = []token{token1, token2}
}
