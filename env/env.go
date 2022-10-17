package env

import (
	. "lisp-go/object"
)

type Env struct {
	Parent *Env
	Vars   map[string]Object
}

func New() *Env {
	return &Env{Vars: make(map[string]Object)}
}

func Extend(parent *Env) *Env {
	newE := New()
	newE.Parent = parent
	return newE
}

func (e Env) Get(key string) (Object, bool) {
	v, ok := e.Vars[key]
	return v, ok
}

func (e *Env) Set(key string, val Object) {
	e.Vars[key] = val
}
