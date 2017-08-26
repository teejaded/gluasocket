package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func eolFn(l *lua.LState) int {
	l.RaiseError("mime.eol(C, D, marker) not implemented yet") // TODO
	return 0
}
