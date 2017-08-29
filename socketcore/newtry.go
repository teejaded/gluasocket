package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func newtryFn(l *lua.LState) int {
	l.RaiseError("socket.newtry(finalizer) not implemented yet")
	return 0
}
