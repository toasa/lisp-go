package object

import (
	"fmt"
	"strings"
)

type ObjectKind int

const (
	Int ObjectKind = iota
	List
	Symbol
)

type Object struct {
	Kind   ObjectKind
	Val    int
	List   []Object
	Symbol string
}

func IntObject(n int) Object {
	return Object{Kind: Int, Val: n}
}

func ListObject(list []Object) Object {
	return Object{Kind: List, List: list}
}

func SymbolObject(s string) Object {
	return Object{Kind: Symbol, Symbol: s}
}

func Equal(o1, o2 Object) bool {
	if o1.Kind != o2.Kind {
		return false
	}
	switch o1.Kind {
	case Int:
		return o1.Val == o2.Val
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
	}

	// (TODO?) Invalid object type, how to handle it?
	return false
}

func (o Object) String() string {
	switch o.Kind {
	case Int:
		return fmt.Sprintf("%d", o.Val)
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
	default:
		return ""
	}
}
