package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func absoluteFn(l *lua.LState) int {
	l.RaiseError("socket.url.absolute(base,relative) not implemented yet") // TODO
	return 0
}
