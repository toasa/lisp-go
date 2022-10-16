package token

import (
	"testing"
)

type test struct {
	input    string
	expected []Token
}

func TestTokenize(t *testing.T) {
	tests := []test{
		{
			input: "(+ 1 2)",
			expected: []Token{
				LParenToken(),
				SymbolToken("+"),
				IntToken(1),
				IntToken(2),
				RParenToken(),
			},
		},
		{
			input: `(
                (define r 10)
                (define pi 314)
                (* pi (* r r))
            )`,
			expected: []Token{
				LParenToken(),
				LParenToken(),
				SymbolToken("define"),
				SymbolToken("r"),
				IntToken(10),
				RParenToken(),
				LParenToken(),
				SymbolToken("define"),
				SymbolToken("pi"),
				IntToken(314),
				RParenToken(),
				LParenToken(),
				SymbolToken("*"),
				SymbolToken("pi"),
				LParenToken(),
				SymbolToken("*"),
				SymbolToken("r"),
				SymbolToken("r"),
				RParenToken(),
				RParenToken(),
				RParenToken(),
			},
		},
	}

	for _, test := range tests {
		tokens := Tokenize(test.input)
		if len(tokens) != len(test.expected) {
			t.Errorf("Token length mismatch (Actual: %d, Expected: %d)", len(tokens), len(test.expected))
		}

		for i, token := range tokens {
			expected := test.expected[i]
			if token.Kind != expected.Kind {
				t.Errorf("Token kind mismatch (Actual: %s, Expected: %s)",
					token, expected)
			}
			if token.Kind == Int && token.Val != expected.Val {
				t.Errorf("Int token value mismatch (Actual: %d, Expected: %d)",
					token.Val, expected.Val)
			}
			if token.Kind == Symbol && token.Symbol != expected.Symbol {
				t.Errorf("Symbol token value mismatch (Actual: %s, Expected: %s)",
					token.Symbol, expected.Symbol)
			}
		}
	}
}
