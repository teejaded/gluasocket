package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func unqpFn(l *lua.LState) int {
	l.RaiseError("mime.unqp(C,D) not implemented yet") // TODO
	return 0
}
