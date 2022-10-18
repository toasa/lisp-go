package main

import (
	"bufio"
	"fmt"
	"lisp-go/env"
	"lisp-go/eval"
	"lisp-go/parse"
	"lisp-go/token"
	"os"
)

func main() {
	PROMPT := "> "

	env := env.New()

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s", PROMPT)

		s.Scan()
		input := s.Text()

		tokens := token.Tokenize(input)
		obj, err := parse.Parse(tokens)
		if err != nil {
			fmt.Printf("Parse failed: %s", err)
			continue
		}

		obj, err = eval.Eval(obj, env)
		if err != nil {
			fmt.Printf("Eval failed: %s", err)
			continue
		}
		fmt.Println(obj)
	}
}
