package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func escapeFn(l *lua.LState) int {
	l.RaiseError("socket.url.escape(content) not implemented yet") // TODO
	return 0
}
