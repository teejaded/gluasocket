package gluasocket_socket

import (
	"github.com/yuin/gopher-lua"
)

func tryFn(l *lua.LState) int {
	l.RaiseError("socket.try(ret, ... ret<n>) not implemented yet")
	return 0
}
