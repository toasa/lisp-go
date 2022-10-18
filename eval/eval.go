package eval

import (
	"fmt"
	. "lisp-go/env"
	. "lisp-go/object"
)

func Eval(obj Object, env *Env) (Object, error) {
	return evalObj(obj, env)
}

func evalObj(obj Object, env *Env) (Object, error) {
	switch obj.Kind {
	case Int:
		return IntObject(obj.Val), nil
	case Void:
		return VoidObject(), nil
	case List:
		return evalList(obj, env)
	case Bool:
		return BoolObject(obj.Bool), nil
	}
	return None, fmt.Errorf("Invalid Object: %s", obj)
}

func evalList(list Object, env *Env) (Object, error) {
	head := list.List[0]

	switch head.Kind {
	case Symbol:
		switch head.Symbol {
		case "+", "-", "*", "/", "<", ">", "=", "!=":
			return evalBinaryOp(list, env)
		case "if":
			return evalIf(list, env)
		}
	default:
		new_list := []Object{}
		for _, elem := range list.List {
			res, _ := evalObj(elem, env)
			if res.Kind == Void {
				continue
			}
			new_list = append(new_list, res)
		}
		return ListObject(new_list), nil
	}

	return None, fmt.Errorf("Failed to eval list")
}

func evalBinaryOp(list Object, env *Env) (Object, error) {
	if len(list.List) != 3 {
		return None, fmt.Errorf("Invalid number of arguments for infix operator")
	}

	op := list.List[0]
	lhs, _ := evalObj(list.List[1], env)
	rhs, _ := evalObj(list.List[2], env)

	if lhs.Kind != Int {
		return None, fmt.Errorf("Left operand must be an integer (%s)", lhs)
	}
	if rhs.Kind != Int {
		return None, fmt.Errorf("Right operand must be an integer (%s)", rhs)
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
			return None, fmt.Errorf("%s unsupported", op.Symbol)
		}
	default:
		return None, fmt.Errorf("Operator must be a symbol")
	}
}

func evalIf(list Object, env *Env) (Object, error) {
	if len(list.List) != 4 {
		return None, fmt.Errorf("Invalid number of arguments for if statement")
	}

	cond_obj, _ := evalObj(list.List[1], env)
	if cond_obj.Kind != Bool {
		return None, fmt.Errorf("Condition must be boolean")
	}
	cond := cond_obj.Bool

	if cond {
		return evalObj(list.List[2], env)
	} else {
		return evalObj(list.List[3], env)
	}
}
