package gluasocket

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"connect": connectFn,
	"gettime": gettimeFn,
	"select":  selectFn,
}

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	registerClientType(l)

	return 1
}

// ----------------------------------------------------------------------------

func registerClientType(l *lua.LState) {
	mt := l.NewTypeMetatable("client")
	l.SetGlobal("client", mt)
	l.SetField(mt, "__index", l.SetFuncs(l.NewTable(), clientMethods))
}
