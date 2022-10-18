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
	case Symbol:
		return evalSymbol(obj, env)
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
		case "define":
			return evalDef(list, env)
		case "lambda":
			return evalFuncDef(list)
		default:
			return evalFuncCall(list, env)
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
}

func evalSymbol(s Object, env *Env) (Object, error) {
	val, ok := env.Get(s.Symbol)
	if !ok {
		return None, fmt.Errorf("Unbound symbol: %s", s)
	}
	return val, nil
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

func evalDef(list Object, env *Env) (Object, error) {
	if len(list.List) != 3 {
		return None, fmt.Errorf("Invalid number of arguments for define")
	}

	key := list.List[1]
	if key.Kind != Symbol {
		return None, fmt.Errorf("Define key must be symbol")
	}
	val, _ := evalObj(list.List[2], env)
	env.Set(key.Symbol, val)
	return VoidObject(), nil
}

func evalFuncDef(list Object) (Object, error) {
	if list.List[1].Kind != List {
		return None, fmt.Errorf("Invalid lambda")
	}

	params := []string{}
	for _, param := range list.List[1].List {
		if param.Kind != Symbol {
			return None, fmt.Errorf("Invalid lambda parameter")
		}
		params = append(params, param.Symbol)
	}

	if list.List[2].Kind != List {
		return None, fmt.Errorf("Invalid lamnda body")
	}
	body := list.List[2].List

	lambda := Object{
		Kind: Lambda,
		Lambda: LambdaObject{
			Params: params,
			Body:   body,
		},
	}

	return lambda, nil
}

func evalFuncCall(list Object, env *Env) (Object, error) {
	fname := list.List[0].Symbol

	f, ok := env.Get(fname)
	if !ok {
		return None, fmt.Errorf("Unbound symbol: %s", fname)
	}

	if f.Kind != Lambda {
		return None, fmt.Errorf("Not a lambda: %s", f)
	}

	newEnv := Extend(env)

	lambda := f.Lambda
	for i, param := range lambda.Params {
		arg, _ := evalObj(list.List[i+1], env)
		newEnv.Set(param, arg)
	}

	return evalObj(ListObject(lambda.Body), newEnv)
}
