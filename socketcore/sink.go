package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func sinkFn(l *lua.LState) int {
	l.RaiseError("socket.sink(mode,socket) not implemented yet")
	return 0
}
