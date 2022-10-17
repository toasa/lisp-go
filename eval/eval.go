package eval

import (
	"fmt"
	. "lisp-go/object"
)

func Eval(obj Object) (Object, error) {
	return evalObj(obj)
}

func evalObj(obj Object) (Object, error) {
	switch obj.Kind {
	case Int:
		return IntObject(obj.Val), nil
	case Void:
		return VoidObject(), nil
	case List:
		return evalList(obj)
	}
	return Object{}, fmt.Errorf("Invalid Object: %s", obj)
}

func evalList(list Object) (Object, error) {
	head := list.List[0]

	switch head.Kind {
	case Symbol:
		switch head.Symbol {
		case "+", "-", "*", "/", "<", ">", "=", "!=":
			return evalBinaryOp(list)
		}
	default:
		return Object{}, fmt.Errorf("Unsupport list %s", list)
	}

	new_list := []Object{}
	for _, elem := range list.List {
		res, _ := evalObj(elem)
		if res.Kind == Void {
			continue
		}
		new_list = append(new_list, res)
	}
	return ListObject(new_list), nil
}

func evalBinaryOp(list Object) (Object, error) {
	if len(list.List) != 3 {
		return Object{}, fmt.Errorf("Invalid number of arguments for infix operator")
	}

	op := list.List[0]
	lhs, _ := evalObj(list.List[1])
	rhs, _ := evalObj(list.List[2])

	if lhs.Kind != Int {
		return Object{}, fmt.Errorf("Left operand must be an integer (%s)", lhs)
	}
	if rhs.Kind != Int {
		return Object{}, fmt.Errorf("Right operand must be an integer (%s)", rhs)
	}

	lval := lhs.Val
	rval := rhs.Val

	switch op.Kind {
	case Symbol:
		switch op.Symbol {
		case "+":
			return IntObject(lval + rval), nil
		case "-":
			return IntObject(lval - rval), nil
		case "*":
			return IntObject(lval * rval), nil
		case "/":
			return IntObject(lval / rval), nil
		case "<":
			return BoolObject(lval < rval), nil
		case ">":
			return BoolObject(lval > rval), nil
		case "=":
			return BoolObject(lval == rval), nil
		case "!=":
			return BoolObject(lval != rval), nil
		default:
			return Object{}, fmt.Errorf("%s unsupported", op.Symbol)
		}
	default:
		return Object{}, fmt.Errorf("Operator must be a symbol")
	}
}
