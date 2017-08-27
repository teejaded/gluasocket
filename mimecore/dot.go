package gluasocket_mimecore

import (
	"github.com/yuin/gopher-lua"
)

func dotFn(l *lua.LState) int {
	l.RaiseError("mime.dot(m, B) not implemented yet") // TODO
	return 0
}
