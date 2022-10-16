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

	list := []Object{}
	t = tokens[1]
	if t.Kind != token.Int {
		return Object{}, fmt.Errorf("List contains Int only, but got %s", t)
	}
	list = append(list, IntObject(t.Val))

	return ListObject(list), nil
}
