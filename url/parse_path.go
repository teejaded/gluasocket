package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func parsePathFn(l *lua.LState) int {
	l.RaiseError("socket.url.parse_path(path) not implemented yet") // TODO
	return 0
}
