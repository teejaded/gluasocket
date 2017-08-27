package gluasocket_socket

import (
	"github.com/yuin/gopher-lua"
)

func skipFn(l *lua.LState) int {
	l.RaiseError("socket.skip(d, [... ret<n>]) not implemented yet")
	return 0
}
