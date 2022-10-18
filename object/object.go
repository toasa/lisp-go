package object

import (
	"fmt"
	"strings"
)

type ObjectKind int

const (
	Int ObjectKind = iota
	Void
	List
	Symbol
	Bool
	Lambda
)

type Object struct {
	Kind   ObjectKind
	Val    int
	List   []Object
	Symbol string
	Bool   bool
	Lambda LambdaObject
}

type LambdaObject struct {
	params []string
	body   []Object
}

var None Object = Object{}

func IntObject(n int) Object {
	return Object{Kind: Int, Val: n}
}

func VoidObject() Object {
	return Object{Kind: Void}
}

func ListObject(list []Object) Object {
	return Object{Kind: List, List: list}
}

func SymbolObject(s string) Object {
	return Object{Kind: Symbol, Symbol: s}
}

func BoolObject(b bool) Object {
	return Object{Kind: Bool, Bool: b}
}

func Equal(o1, o2 Object) bool {
	if o1.Kind != o2.Kind {
		return false
	}
	switch o1.Kind {
	case Int:
		return o1.Val == o2.Val
	case Void:
		return true
	case List:
		if len(o1.List) != len(o2.List) {
			return false
		}
		for i, e1 := range o1.List {
			e2 := o2.List[i]
			if !Equal(e1, e2) {
				return false
			}
		}
		return true
	case Symbol:
		return o1.Symbol == o2.Symbol
	case Bool:
		return o1.Bool == o2.Bool
	case Lambda:
		if len(o1.Lambda.params) != len(o2.Lambda.params) {
			return false
		}
		for i, param1 := range o1.Lambda.params {
			param2 := o2.Lambda.params[i]
			if param1 != param2 {
				return false
			}
		}

		if len(o1.Lambda.body) != len(o2.Lambda.body) {
			return false
		}
		for i, expr1 := range o1.Lambda.body {
			expr2 := o2.Lambda.body[i]
			if !Equal(expr1, expr2) {
				return false
			}
		}

		return true
	}

	// (TODO?) Invalid object type, how to handle it?
	return false
}

func (o Object) String() string {
	switch o.Kind {
	case Int:
		return fmt.Sprintf("%d", o.Val)
	case Void:
		return "Void"
	case List: // List
		var l strings.Builder

		l.WriteString("(")
		for i, e := range o.List {
			if i > 0 {
				l.WriteString(" ")
			}
			l.WriteString(e.String())
		}
		l.WriteString(")")

		return l.String()
	case Symbol:
		return o.Symbol
	case Bool:
		return fmt.Sprintf("%v", o.Bool)
	case Lambda:
		return o.Lambda.String()
	default:
		return ""
	}
}

func (l LambdaObject) String() string {
	var b strings.Builder
	b.WriteString("Lambda(")
	for _, param := range l.params {
		b.WriteString(param)
		b.WriteString(" ")
	}
	b.WriteString(")")

	for _, expr := range l.body {
		b.WriteString(expr.String())
	}

	return b.String()
}
