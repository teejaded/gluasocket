package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func unescapeFn(l *lua.LState) int {
	l.RaiseError("socket.url.unescape(content) not implemented yet") // TODO
	return 0
}
