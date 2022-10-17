package parse

import (
	"fmt"
	. "lisp-go/object"
	"lisp-go/token"
)

func Parse(tokens []token.Token) (Object, error) {
	return parseList(tokens)
}

func parseList(tokens []token.Token) (Object, error) {
	t := tokens[0]
	if t.Kind != token.LParen {
		return Object{}, fmt.Errorf("Expected LParen, got %s", t)
	}
	tokens = tokens[1:]

	list := []Object{}
	for len(tokens) > 0 {
		t = tokens[0]
		tokens = tokens[1:]

		switch t.Kind {
		case token.Int:
			list = append(list, IntObject(t.Val))
		case token.Symbol:
			list = append(list, SymbolObject(t.Symbol))
		case token.RParen:
			return ListObject(list), nil
		default:
			return Object{}, fmt.Errorf("Invalid token %s", t)
		}
	}

	return ListObject(list), nil
}