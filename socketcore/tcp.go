package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func tcpFn(l *lua.LState) int {
	l.RaiseError("socket.tcp() not implemented yet")
	return 0
}
