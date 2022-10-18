package eval

import (
	"lisp-go/env"
	. "lisp-go/object"
	"lisp-go/parse"
	"lisp-go/token"
	"testing"
)

type test struct {
	input    string
	expected Object
}

var testEnv *env.Env = env.New()

func TestEval(t *testing.T) {
	tests := []test{
		{
			input:    "(+ 10 20)",
			expected: IntObject(30),
		},
		{
			input:    "(+ (+ 3 4) 5)",
			expected: IntObject(12),
		},
		{
			input:    "(+ 2 (* 3 4))",
			expected: IntObject(14),
		},
		{
			input:    "(+ (* 2 3) 4)",
			expected: IntObject(10),
		},
		{
			input:    "(< 3 4)",
			expected: BoolObject(true),
		},
		{
			input:    "(= 3 4)",
			expected: BoolObject(false),
		},
		{
			input:    "(if (< 3 4) 10 20)",
			expected: IntObject(10),
		},
		{
			input:    "(if (> 3 4) 10 20)",
			expected: IntObject(20),
		},
		{
			input: `(
				(define x 11)
				(* x x))
				)`,
			expected: ListObject(
				[]Object{
					IntObject(121),
				},
			),
		},
		{
			input: `(
				(define sqr (lambda (r) (* r r))) 
				(sqr 11)
			   )`,
			expected: ListObject(
				[]Object{
					IntObject(121),
				},
			),
		},
		{
			input: `(
				(define f (lambda (a b c) (+ (* a b) c)))
				(f 2 3 4)
			   )`,
			expected: ListObject(
				[]Object{
					IntObject(10),
				},
			),
		},
		{
			input: `(
		              (define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))
		              (fib 10)
		          )`,
			expected: ListObject(
				[]Object{
					IntObject(89),
				},
			),
		},
		{
			input: `(
                (define fact (lambda (n) (if (< n 1) 1 (* n (fact (- n 1))))))
                (fact 5)
            )`,
			expected: ListObject(
				[]Object{
					IntObject(5 * 4 * 3 * 2 * 1),
				},
			),
		},
		{
			input: `(
                (define pi 314)
                (define r 10)
                (define sqr (lambda (r) (* r r)))
                (define area (lambda (r) (* pi (sqr r))))
                (area r)
            )`,
			expected: ListObject(
				[]Object{
					IntObject(314 * 10 * 10),
				},
			),
		},
	}

	for _, test := range tests {
		tokens := token.Tokenize(test.input)
		obj, err := parse.Parse(tokens)
		if err != nil {
			t.Errorf("Parse failed: %s", err)
		}

		list, err := Eval(obj, testEnv)
		if err != nil {
			t.Errorf("Eval failed: %s", err)
		}

		expected := test.expected
		if !Equal(list, expected) {
			t.Errorf("Object mismatch (Actual: %s, Expected: %s)", list, expected)
		}
	}
}
