package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func udpFn(l *lua.LState) int {
	l.RaiseError("socket.udp() not implemented yet")
	return 0
}
