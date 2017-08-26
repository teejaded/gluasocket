package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func wrpFn(l *lua.LState) int {
	l.RaiseError("mime.wrp(n,B,length) not implemented yet") // TODO
	return 0
}
