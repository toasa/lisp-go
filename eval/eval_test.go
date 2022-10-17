package eval

import (
	. "lisp-go/object"
	"testing"
)

type test struct {
	input    Object
	expected Object
}

func TestEval(t *testing.T) {
	tests := []test{
		{
			input: ListObject(
				[]Object{
					SymbolObject("+"),
					IntObject(10),
					IntObject(20),
				}),
			expected: IntObject(30),
		},
		{
			input: ListObject(
				[]Object{
					SymbolObject("+"),
					ListObject(
						[]Object{
							SymbolObject("+"),
							IntObject(3),
							IntObject(4),
						}),
					IntObject(5),
				}),
			expected: IntObject(12),
		},
		{
			input: ListObject(
				[]Object{
					SymbolObject("+"),
					IntObject(2),
					ListObject(
						[]Object{
							SymbolObject("*"),
							IntObject(3),
							IntObject(4),
						}),
				}),
			expected: IntObject(14),
		},
		{
			input: ListObject(
				[]Object{
					SymbolObject("+"),
					ListObject(
						[]Object{
							SymbolObject("*"),
							IntObject(2),
							IntObject(3),
						}),
					IntObject(4),
				}),
			expected: IntObject(10),
		},
		{
			input: ListObject(
				[]Object{
					SymbolObject("<"),
					IntObject(3),
					IntObject(4),
				}),
			expected: BoolObject(true),
		},
		{
			input: ListObject(
				[]Object{
					SymbolObject("="),
					IntObject(3),
					IntObject(4),
				}),
			expected: BoolObject(false),
		},
	}

	for _, test := range tests {
		list, err := Eval(test.input)
		if err != nil {
			t.Errorf("Eval failed: %s", err)
		}

		expected := test.expected
		if !Equal(list, expected) {
			t.Errorf("Object mismatch (Actual: %s, Expected: %s)", list, expected)
		}
	}
}
