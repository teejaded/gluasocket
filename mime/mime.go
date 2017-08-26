package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	// high-level filters
	"decode":    decodeFn,
	"encode":    encodeFn,
	"normalize": normalizeFn,
	"stuff":     stuffFn,
	"wrap":      wrapFn,

	// low-level filters
	"b64":   b64Fn,
	"dot":   dotFn,
	"eol":   eolFn,
	"qp":    qpFn,
	"qpwrp": qpwrpFn,
	"unb64": unb64Fn,
	"unqp":  unqpFn,
	"wrp":   wrpFn,
}

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	return 1
}
