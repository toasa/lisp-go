package token

import (
	"fmt"
	"strconv"
	"strings"
)

type TokenKind int

const (
	Int TokenKind = iota
	Symbol
	LParen
	RParen
)

type Token struct {
	Kind   TokenKind
	Val    int    // If kind is Int, its value
	Symbol string // If kind is Symbol, its value
}

func IntToken(n int) Token {
	return Token{Kind: Int, Val: n}
}

func LParenToken() Token {
	return Token{Kind: LParen}
}

func RParenToken() Token {
	return Token{Kind: RParen}
}

func SymbolToken(s string) Token {
	return Token{Kind: Symbol, Symbol: s}
}

func Tokenize(program string) []Token {
	program = strings.ReplaceAll(program, "(", " ( ")
	program = strings.ReplaceAll(program, ")", " ) ")
	words := strings.Fields(program)

	tokens := []Token{}
	for _, word := range words {
		switch word {
		case "(":
			tokens = append(tokens, LParenToken())
		case ")":
			tokens = append(tokens, RParenToken())
		default:
			if n, err := strconv.Atoi(word); err == nil {
				tokens = append(tokens, IntToken(n))
			} else {
				tokens = append(tokens, SymbolToken(word))
			}
		}
	}

	return tokens
}

func (t Token) String() string {
	switch t.Kind {
	case Int:
		return fmt.Sprintf("Int(%d)", t.Val)
	case Symbol:
		return fmt.Sprintf("Symbol(%s)", t.Symbol)
	case LParen:
		return "LParen"
	default:
		return "RParen"
	}
}
