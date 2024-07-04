package command

import (
	"fmt"

	"github.com/korepanov/gopy/internal/cerrors"
)

// looks ahead to the next shortest token but does not remove the token from subinput
func (c *Command) lookAhead() (Token, error) {
	var buf string

	for _, ch := range c.subinput {
		buf += string(ch)
		bufType := Dictionary().Find(lexeme(buf))

		if bufType != 0 {
			return Token{Lex: lexeme(buf), T: bufType}, nil
		}
	}

	if len(c.subinput) == 0 {
		return Token{}, nil
	}

	return Token{Lex: lexeme(buf), T: 0}, cerrors.ErrNoToken
}

// gets next longest token and removes this token from subinput
func (c *Command) nextToken() (Token, error) {

	newToken, err := c.lookAhead()

	for err != nil {
		return newToken, err
	}

	c.subinput = c.subinput[len(newToken.Lex):]

	if Dictionary().IsStop(newToken.T) {
		return newToken, nil
	}

	for aheadToken, err := c.lookAhead(); (err == cerrors.ErrNoToken || !Dictionary().IsStop(aheadToken.T)) &&
		len(aheadToken.Lex) != 0; aheadToken, err = c.lookAhead() {

		buf := string(newToken.Lex)

		for idx, ch := range c.subinput {
			buf += string(ch)
			bufType := Dictionary().Find(lexeme(buf))

			if bufType != 0 {
				newToken.Lex = lexeme(buf)
				newToken.T = bufType
				c.subinput = c.subinput[len(newToken.Lex):]
				break
			}

			if idx == len(c.subinput)-1 {
				return Token{}, cerrors.ErrNoToken
			}
		}
	}

	return newToken, nil

}

func (c *Command) LexicalAnalyze() error {

	c.subinput = c.Input

	for newToken, err := c.nextToken(); len(newToken.Lex) != 0; newToken, err = c.nextToken() {

		if err != nil {
			return fmt.Errorf("%s : %s", cerrors.ErrLexAnalysis, err)
		}

		c.Tokens = append(c.Tokens, newToken)
	}

	return nil
}
