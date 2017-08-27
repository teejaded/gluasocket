package gluasocket_socket

import (
	"github.com/yuin/gopher-lua"
)

func sourceFn(l *lua.LState) int {
	l.RaiseError("socket.source(mode,socket,length) not implemented yet")
	return 0
}
