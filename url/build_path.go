package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func buildPathFn(l *lua.LState) int {
	l.RaiseError("socket.url.build_path(segments,unsafe) not implemented yet") // TODO
	return 0
}
