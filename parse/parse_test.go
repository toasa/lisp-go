package parse

import (
	. "lisp-go/object"
	"lisp-go/token"
	"testing"
)

type test struct {
	input    []token.Token
	expected Object
}

func TestParse(t *testing.T) {
	tests := []test{
		{
			input: []token.Token{
				token.LParenToken(),
				token.IntToken(10),
				token.RParenToken(),
			},
			expected: Object{
				Kind: List,
				List: []Object{
					IntObject(10),
				},
			},
		},
	}

	for _, test := range tests {
		list, err := Parse(test.input)
		if err != nil {
			t.Errorf("Parse failed: %s", err)
		}

		expected := test.expected
		if !Equal(list, expected) {
			t.Errorf("Object mismatch (Actual: %s, Expected: %s)", list, expected)
		}
	}
}
