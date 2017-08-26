package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func unb64Fn(l *lua.LState) int {
	l.RaiseError("mime.unb64(C,D) not implemented yet") // TODO
	return 0
}
