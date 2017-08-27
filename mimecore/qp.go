package gluasocket_mimecore

import (
	"github.com/yuin/gopher-lua"
)

func qpFn(l *lua.LState) int {
	l.RaiseError("mime.qp(C,D,marker) not implemented yet") // TODO
	return 0
}
