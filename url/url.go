package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"absolute":   absoluteFn,
	"build":      buildFn,
	"build_path": buildPathFn,
	"escape":     escapeFn,
	"parse":      parseFn,
	"parse_path": parsePathFn,
	"unescape":   unescapeFn,
}

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	return 1
}
