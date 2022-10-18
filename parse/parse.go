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
	head := tokens[0]
	if head.Kind != token.LParen {
		return None, fmt.Errorf("Expected LParen, got %s", head)
	}
	tokens = tokens[1:]

	list := []Object{}
	for len(tokens) > 0 {
		head = tokens[0]
		tokens = tokens[1:]

		switch head.Kind {
		case token.Int:
			list = append(list, IntObject(head.Val))
		case token.Symbol:
			list = append(list, SymbolObject(head.Symbol))
		case token.LParen:
			tokens = append([]token.Token{head}, tokens...)
			sub_list, _ := parseList(tokens)
			list = append(list, sub_list)
			tokens = tmp
		case token.RParen:
			tmp = tokens
			return ListObject(list), nil
		default:
			return None, fmt.Errorf("Invalid token %s", head)
		}
	}

	return ListObject(list), nil
}
