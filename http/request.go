package gluasocket_http

import (
	"github.com/yuin/gopher-lua"
)

func requestFn(l *lua.LState) int {
	l.RaiseError("socket.http.request(url, body) not implemented yet") // TODO
	return 0
}
