package internal

import (
	"fmt"
)

// looks ahead to the next shortest token but does not remove the token from input
func (c *command) lookAhead() (token, error) {
	var buf string

	for _, ch := range c.input {
		buf += string(ch)
		bufType := dictionary().Find(lexeme(buf))

		if bufType != 0 {
			return token{lex: lexeme(buf), t: bufType}, nil
		}
	}

	if len(c.input) == 0 {
		return token{}, nil
	}

	return token{lex: lexeme(buf), t: 0}, ErrNoToken
}

// gets next longest token and removes this token from input
func (c *command) nextToken() (token, error) {

	newToken, err := c.lookAhead()

	for err != nil {
		return newToken, err
	}

	c.input = c.input[len(newToken.lex):]

	if dictionary().IsStop(newToken.t) {
		return newToken, nil
	}

	for aheadToken, err := c.lookAhead(); (err == ErrNoToken || !dictionary().IsStop(aheadToken.t)) &&
		len(aheadToken.lex) != 0; aheadToken, err = c.lookAhead() {

		buf := string(newToken.lex)

		for idx, ch := range c.input {
			buf += string(ch)
			bufType := dictionary().Find(lexeme(buf))

			if bufType != 0 {
				newToken.lex = lexeme(buf)
				c.input = c.input[len(newToken.lex):]
				break
			}

			if idx == len(c.input)-1 {
				return token{}, ErrNoToken
			}
		}
	}

	return newToken, nil

}

func (c *command) lexicalAnalyze() error {

	var err error

	for newToken, err := c.nextToken(); err == nil && len(newToken.lex) != 0; newToken, err = c.nextToken() {
		c.tokens = append(c.tokens, newToken)
	}

	if err != nil {
		return fmt.Errorf("%q : %q", ErrLexAnalysis, ErrNoToken)
	}

	return nil
}
