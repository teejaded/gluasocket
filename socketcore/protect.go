package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func protectFn(l *lua.LState) int {
	l.RaiseError("socket.protect(finalizer) not implemented yet")
	return 0
}
