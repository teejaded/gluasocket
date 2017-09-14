package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func skipFn(L *lua.LState) int {
	L.RaiseError("socket.skip(d, [... ret<n>]) not implemented yet")
	return 0
}
