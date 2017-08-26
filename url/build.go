package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

func buildFn(l *lua.LState) int {
	l.RaiseError("socket.url.build(parsed_url) not implemented yet") // TODO
	return 0
}
