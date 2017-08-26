package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func stuffFn(l *lua.LState) int {
	l.RaiseError("mime.stuff() not implemented yet") // TODO
	return 0
}
