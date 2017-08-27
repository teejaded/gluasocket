package gluasocket_mimecore

import (
	"github.com/yuin/gopher-lua"
)

func b64Fn(l *lua.LState) int {
	l.RaiseError("mime.b64(C, D) not implemented yet") // TODO
	return 0
}
