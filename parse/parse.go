package parse

import (
	"fmt"
	. "lisp-go/object"
	"lisp-go/token"
)

func Parse(tokens []token.Token) (Object, error) {
	return parseList(tokens)
}

// Preserve the tokens sequence that have been read when
// parsing a nested list.
var tmp []token.Token

func parseList(tokens []token.Token) (Object, error) {

	t := tokens[0]
	if t.Kind != token.LParen {
		return None, fmt.Errorf("Expected LParen, got %s", t)
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
		case token.LParen:
			tokens = append([]token.Token{t}, tokens...)
			sub_list, _ := parseList(tokens)
			list = append(list, sub_list)
			tokens = tmp
		case token.RParen:
			tmp = tokens
			return ListObject(list), nil
		default:
			return None, fmt.Errorf("Invalid token %s", t)
		}
	}

	return ListObject(list), nil
}
